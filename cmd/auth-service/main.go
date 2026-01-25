package main

import (
	"context"
	"database/sql"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	auth "github.com/enyaaad/CryptoWalletBackend/api/proto"
	grpcD "github.com/enyaaad/CryptoWalletBackend/internal/auth/delivery/grpc"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/infrastructure/jwt"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/infrastructure/password"
	"github.com/enyaaad/CryptoWalletBackend/internal/auth/infrastructure/postgres"
	authS "github.com/enyaaad/CryptoWalletBackend/internal/auth/infrastructure/service"
	"github.com/enyaaad/CryptoWalletBackend/pkg/logger"
	_ "github.com/lib/pq"
	grpc "google.golang.org/grpc"
)

func main() {
	logger := logger.InitLogger("auth-service")
	logger.Info().Msg("Starting Auth Service")

	dbURL := getEnv("DATABASE_URL", "postgres://cryptowallet:cryptowallet123@localhost:5432/cryptowallet?sslmode=disable")
	jwtSecret := getEnv("JWT_SECRET", "jwtSecret")
	port := getEnv("GRPC_PORT", "50051")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to open DB connection")
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Failed to ping DB")
	}
	logger.Info().Msg("Database connection established")

	userRepo := postgres.NewUserRepository(db)
	passwordHasher := password.NewBcryptHasher()
	jwtSvc := jwt.NewJWTService(jwtSecret, 16*time.Minute, 7*24*time.Hour)
	authSvc := authS.NewAuthService(userRepo, passwordHasher, jwtSvc)
	grpcServer := grpc.NewServer()
	authGRPCHandler := grpcD.NewAuthGRPCServer(authSvc, logger)
	auth.RegisterAuthServiceServer(grpcServer, authGRPCHandler)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to listen gRPC")
	}

	go func() {
		logger.Info().Str("port", port).Msg("gRPC server started")
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal().Err(err).Msg("Failed to serve gRPC")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info().Msg("Shutting down gRPC server")

	stopped := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		logger.Info().Msg("gRPC server stopped gracefully")
	case <-time.After(10 * time.Second):
		logger.Warn().Msg("Graceful shutdown timeout, forcing stop")
		grpcServer.Stop()
	}

	logger.Info().Msg("Auth Service stopped")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
