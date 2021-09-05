package server

import (
	"github.com/gin-gonic/gin"
	"go-notify/controllers"
	"go-notify/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(middlewares.LogFormatter))
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	return router
}
