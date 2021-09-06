package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mw3tv123/go-notify/controllers"
	"github.com/mw3tv123/go-notify/middlewares"
)

// NewRouter return a Go Gin router with all routes be defined
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(middlewares.LogFormatter))
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)
	api := router.Group("/api")
	{
		/*** START MS TEAMS ***/
		msTeamsController := controllers.NewMSTeamsController()
		api.POST("/ms_teams", msTeamsController.Notify)
	}

	return router
}
