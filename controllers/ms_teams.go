package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mw3tv123/go-notify/config"
	"github.com/mw3tv123/go-notify/forms"
	"github.com/mw3tv123/go-notify/models"
)

// MSTeamController ...
type MSTeamController struct {
	msTeamsService *models.MSTeamsService
}

var msTeamsForm = new(forms.MSTeamsForm)

// NewMSTeamsController Initiate MSTeams service for transmitting message to MS Teams
func NewMSTeamsController() MSTeamController {
	msTeamsController := MSTeamController{
		msTeamsService: models.NewMSTeamsService(),
	}
	msTeamsController.msTeamsService.AddReceivers(config.GetConfig("MS_TEAMS_WEBHOOK"))

	return msTeamsController
}

// Notify send a simple notification to MS Teams webhook
func (ms MSTeamController) Notify(c *gin.Context) {
	var msTeamsNotifyForm forms.CreateMSTeamsNotifyForm

	if validationErr := c.ShouldBindJSON(&msTeamsNotifyForm); validationErr != nil {
		message := msTeamsForm.CreateNotify(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err := ms.msTeamsService.SendMessage(context.Background(), msTeamsNotifyForm.Title, msTeamsNotifyForm.Content)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified"})
}

// Alert create an alert card and then send it to all MS Teams webhook
func (ms MSTeamController) Alert(c *gin.Context) {
	var msTeamsAlertForm forms.CreateMSTeamsAlertForm

	if validationErr := c.ShouldBindJSON(&msTeamsAlertForm); validationErr != nil {
		message := msTeamsForm.CreateAlert(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	msg, err := ms.msTeamsService.GenerateAlertCard(msTeamsAlertForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	err = ms.msTeamsService.Send(context.Background(), *msg)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified"})
}
