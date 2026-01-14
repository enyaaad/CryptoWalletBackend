package main

import (
	"github.com/enyaaad/CryptoWalletBackend/internal/api-gateway/middleware"
	"github.com/enyaaad/CryptoWalletBackend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := logger.InitLogger("api-gateway")
	logger.Info().Msg("Starting APIGATEWAY")

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "all good!"})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Use(middleware.MetricsMiddleware())
	r.Use(middleware.LoggerNiddleware(logger))

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register")
			auth.POST("/login")
		}

		protected := v1.Group("")
		{
			wallets := protected.Group("/wallets")
			{
				wallets.POST("/get")
			}
		}
	}

	r.Run(":8080")
}
