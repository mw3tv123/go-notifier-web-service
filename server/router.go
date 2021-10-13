package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mw3tv123/go-notify/controllers"
	"github.com/mw3tv123/go-notify/forms"
	"github.com/mw3tv123/go-notify/middlewares"
)

// NewRouter return a Go Gin router with all routes be defined
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	// Custom form log formatter
	router.Use(gin.LoggerWithFormatter(middlewares.LogFormatter))

	// Custom form validator
	binding.Validator = new(forms.DefaultValidator)

	/*** HEALTH CHECK API ***/
	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	/*** NOTIFY GROUP API ***/
	notificationController := controllers.NewNotificationController()
	router.POST("/notify", notificationController.Notify)

	return router
}
