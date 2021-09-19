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
	msTeamsController := controllers.NewMSTeamsController()

	/*** NOTIFY GROUP API ***/
	notifyApi := router.Group("/notify")
	{
		notifyApi.POST("/ms_teams", msTeamsController.Notify)
	}
	/*** ALERT GROUP API ***/
	alertApi := router.Group("/alert")
	{
		alertApi.POST("/ms_teams", msTeamsController.Alert)
	}

	return router
}
