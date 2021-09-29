package forms

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

// SupportedChannels represents list of current supported channels
var SupportedChannels = []string{"teams"}

// RequestForm ...
type RequestForm struct{}

// RequestMessageForm represents a request form client for sending out a message to target channel.
type RequestMessageForm struct {
	// Title is the title property of a card. is meant to be rendered in a
	// prominent way, at the very top of the card. Use it to introduce the
	// content of the card in such a way users will immediately know what to
	// expect.
	Title string `form:"title" json:"title" binding:"required,min=3,max=100"`
	// Content is the message which send to target channels. Supported only short message with maximum of 200 characters.
	Content string `form:"content" json:"content" binding:"required,min=3,max=200"`
	// Channels is list of channels client to sent notification to. List of supported channels must add to custom validator.
	Channels []string `form:"channels" json:"channels,omitempty" binding:"dive,supportChannel"`
}

// RequestAlertForm represents a request form client for sending out a alert message card to target channel.
type RequestAlertForm struct {
	// Title is same as Title in RequestMessageForm
	Title string `form:"title" json:"title" binding:"required,min=3,max=100"`

	// Priority or Critical indicate the importance of this message. Priority level in some services may in scale of 5,
	// while other services may have scale up to 10.
	Priority int `form:"priority" json:"priority" binding:"required,min=0,max=10"`

	// ServiceName is the name of service has alert.
	ServiceName string `form:"service_name" json:"service_name" binding:"required,min=3,max=100"`

	// Description describe the information about alert. Maybe detail, maybe summary.
	Description string `form:"description" json:"description,omitempty" binding:"max=100"`

	// CreateDate is the date the service received alert. If left empty, system will choose
	// current timestamp as default CreateDate to sent to channel.
	CreateDate time.Time `form:"create_date" json:"create_date,omitempty"`

	// Channels same as channels in normal message. Reference to RequestMessageForm Channels.
	Channels []string `form:"channels" json:"channels,omitempty" binding:"dive,supportChannel"`
}

// evaluateErrorMessage evaluate error message by combine tag, field name and param
func (f RequestForm) evaluateErrorMessage(tag, fieldName, param string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return fmt.Sprintf("Please enter %s", fieldName)
		}
		return errMsg[0]
	case "min", "max":
		compareString := "greater"
		if tag == "min" {
			compareString = "lower"
		}
		return fmt.Sprintf("%s should be %s than %s", fieldName, compareString, param)
	case "supportChannel":
		return "Channel is not support or name is wrong."
	default:
		return "Something went wrong, please try again later"
	}
}

// GetValidatedErrorMessage will return error message when form not pass validation
func (f RequestForm) GetValidatedErrorMessage(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			return f.evaluateErrorMessage(err.Tag(), err.Field(), err.Param())
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
