package router

import (
	"schedule-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handler.HealthCheck)

	return r
}
