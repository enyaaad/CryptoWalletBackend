package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"

	httpHandlers "github.com/enyaaad/CryptoWalletBackend/internal/api-gateway/delivery/http"
	grpcClient "github.com/enyaaad/CryptoWalletBackend/internal/api-gateway/infrastructure/grpc"
	"github.com/enyaaad/CryptoWalletBackend/internal/api-gateway/middleware"
	"github.com/enyaaad/CryptoWalletBackend/pkg/logger"
)

func main() {
	logger := logger.InitLogger("api-gateway")
	logger.Info().Msg("Starting API Gateway")

	// Конфигурация из env
	httpPort := getEnv("HTTP_PORT", "8080")
	authServiceAddr := getEnv("AUTH_SERVICE_ADDR", "auth-service:50051")

	authClient, err := grpcClient.NewAuthClient(authServiceAddr)
	if err != nil {
		logger.Fatal().Err(err).Str("addr", authServiceAddr).Msg("Failed to create auth gRPC client")
	}
	defer authClient.Close()
	logger.Info().Str("addr", authServiceAddr).Msg("Auth gRPC client connected")

	authHandler := httpHandlers.NewAuthHandler(authClient)

	router := setupRouter(logger, authHandler)

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: router,
	}

	go func() {
		logger.Info().Str("port", httpPort).Msg("HTTP server starting")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("failed to start http server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info().Msg("Shutting down API Gateway")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("Server forced to shutdown")
	}

	logger.Info().Msg("API Gateway stopped")
}

func setupRouter(log zerolog.Logger, authHandler *httpHandlers.AuthHandler) *gin.Engine {
	mode := getEnv("MODE", "prod")

	router := gin.New()
	gin.SetMode(mode)

	router.Use(middleware.LoggerNiddleware(log))
	router.Use(middleware.MetricsMiddleware())

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("metrics", gin.WrapH(promhttp.Handler()))

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.POST("/validate", authHandler.ValidateToken)
		}
	}

	return router
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
