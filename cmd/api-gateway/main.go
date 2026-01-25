package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/enyaaad/CryptoWalletBackend/pkg/logger"
)

func main() {
	logger := logger.InitLogger("api-gateway")
	logger.Info().Msg("Starting API Gateway")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
