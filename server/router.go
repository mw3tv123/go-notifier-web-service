package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mw3tv123/go-notify/controllers"
	"github.com/mw3tv123/go-notify/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(middlewares.LogFormatter))
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	return router
}
