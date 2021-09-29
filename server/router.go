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

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	/*** START CONTROLLER ***/
	notificationController := controllers.NewNotificationController()

	/*** NOTIFY GROUP API ***/
	notifyAPI := router.Group("/message")
	{
		notifyAPI.POST("/ms_teams", notificationController.Message)
	}
	/*** ALERT GROUP API ***/
	alertAPI := router.Group("/alert")
	{
		alertAPI.POST("/ms_teams", notificationController.Alert)
	}

	return router
}
