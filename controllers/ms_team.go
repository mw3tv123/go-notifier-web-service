package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mw3tv123/go-notify/config"
	"github.com/mw3tv123/go-notify/forms"
	"github.com/nikoksr/notify/service/msteams"
)

// MSTeamController ...
type MSTeamController struct {
	MsTeamsService *msteams.MSTeams
}

var msTeamsNotifyForm = new(forms.MSTeamsNotifyForm)

// NewMSTeamsController Initiate MSTeams service for transmitting message to MS Teams
func NewMSTeamsController() MSTeamController {
	msTeamsController := MSTeamController{
		MsTeamsService: msteams.New(),
	}
	msTeamsController.MsTeamsService.AddReceivers(config.GetConfig("MS_TEAMS_WEBHOOK"))

	return msTeamsController
}

// Notify ...
func (ms MSTeamController) Notify(c *gin.Context) {
	var msTeamNotifyForm forms.CreateMSTeamNotifyForm

	if validationErr := c.ShouldBindJSON(&msTeamNotifyForm); validationErr != nil {
		message := msTeamsNotifyForm.CreateNotify(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err := ms.MsTeamsService.Send(context.Background(), msTeamNotifyForm.Title, msTeamNotifyForm.Content)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified"})
}
