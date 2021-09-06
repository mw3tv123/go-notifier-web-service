package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mw3tv123/go-notify/forms"
	"github.com/nikoksr/notify/service/msteams"
)

// MSTeamController ...
type MSTeamController struct {
	MsTeamsService *msteams.MSTeams
}

var notifyForm = new(forms.NotifyForm)

func NewMSTeamsController() MSTeamController {
	msTeamsController := MSTeamController{
		MsTeamsService: msteams.New(),
	}
	msTeamsController.MsTeamsService.AddReceivers("https://vngms.webhook.office.com/webhookb2/70b30a98-196a-43e7-ae1e-cb67b67e3c0f@7c112a6e-10e2-4e09-afc4-2e37bc60d821/IncomingWebhook/eba3b645a7114041b2a9b3c1d4afc87c/8af52f50-5b7f-4691-8c93-7f54240e8a0f")

	return msTeamsController
}

func (ms MSTeamController) Notify(c *gin.Context) {
	var msTeamNotifyForm forms.MSTeamNotifyForm

	if validationErr := c.ShouldBindJSON(&msTeamNotifyForm); validationErr != nil {
		message := notifyForm.Notify(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err := ms.MsTeamsService.Send(context.Background(), msTeamNotifyForm.Title, msTeamNotifyForm.Message)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified"})
}
