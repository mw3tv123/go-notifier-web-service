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
		},
	}

	return notificationController
}

// Message send a simple message to MS Teams webhook
func (n NotificationController) Message(c *gin.Context) {
	var messageForm forms.RequestMessageForm

	if validationErr := c.ShouldBindJSON(&messageForm); validationErr != nil {
		message := requestForm.GetValidatedErrorMessage(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	if len(messageForm.Channels) <= 0 {
		messageForm.Channels = forms.SupportedChannels
	}

	for _, name := range messageForm.Channels {
		err := n.channels[name].SendMessage(context.Background(), messageForm)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified"})
}

// Alert create an alert card and then send it to all MS Teams webhook
func (n NotificationController) Alert(c *gin.Context) {
	var alertForm forms.RequestAlertForm

	if validationErr := c.ShouldBindJSON(&alertForm); validationErr != nil {
		message := requestForm.GetValidatedErrorMessage(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	if len(alertForm.Channels) <= 0 {
		alertForm.Channels = forms.SupportedChannels
	}

	for _, name := range alertForm.Channels {
		err := n.channels[name].SendAlert(context.Background(), alertForm)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully notified"})
}
