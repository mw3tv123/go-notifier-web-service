package forms

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// SupportedChannels represents list of current supported channels
var SupportedChannels = []string{"teams"}

// RequestForm ...
type RequestForm struct{}

// MessageForm represents a request form client for sending out a message to target channel.
type MessageForm struct {
	// Contents is the message which send to target channels.
	Contents map[string]interface{} `form:"contents" json:"contents" binding:"required"`
	// Channels is list of channels client to sent notification to. List of supported channels must add to custom validator.
	Channels []string `form:"channels" json:"channels,omitempty" binding:"dive,supportChannel"`
	// Template is a name of a template client wish their message to be loaded into.
	Template string `form:"template" json:"template,omitempty"`
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
