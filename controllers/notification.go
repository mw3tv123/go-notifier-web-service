package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mw3tv123/go-notify/config"
	"github.com/mw3tv123/go-notify/forms"
	"github.com/mw3tv123/go-notify/models"
)

// NotificationController ...
type NotificationController struct {
	channels map[string]models.Channel
}

var requestForm = new(forms.RequestForm)

// NewNotificationController Initiate MSTeams service for transmitting message to MS Teams
func NewNotificationController() NotificationController {
	notificationController := NotificationController{
		channels: map[string]models.Channel{
			"teams": models.NewMSTeamsService(config.GetConfig("MS_TEAMS_WEBHOOK")),
			// "email": models.NewEmailService(config.GetConfig("SENDER_ADDRESS"), config.GetConfig("SMTP_HOST_ADDRESS")),
		},
	}

	return notificationController
}

// Notify parse the request form info proper message struct and send to each channel depend on request.
func (n NotificationController) Notify(c *gin.Context) {
	var messageForm forms.MessageForm

	if validationErr := c.ShouldBindJSON(&messageForm); validationErr != nil {
		message := requestForm.GetValidatedErrorMessage(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	if len(messageForm.Channels) <= 0 {
		messageForm.Channels = forms.SupportedChannels
	}

	for _, name := range messageForm.Channels {
		err := n.channels[name].Send(context.Background(), messageForm)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified"})
}
