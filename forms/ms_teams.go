package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// NotifyForm ...
type NotifyForm struct{}

// MSTeamNotifyForm ...
type MSTeamNotifyForm struct {
	Title   string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Message string `form:"message" json:"message" binding:"required,max=200,message"`
}

// Title ...
func (f NotifyForm) Title(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter notify title"
		}
		return errMsg[0]
	case "min", "max":
		return "Notify title should be between 3 to 100 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// Message ...
func (f NotifyForm) Message(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your message"
		}
		return errMsg[0]
	case "min", "max", "message":
		return "Please enter a valid message"
	default:
		return "Something went wrong, please try again later"
	}
}

// Notify ...
func (f NotifyForm) Notify(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Title" {
				return f.Title(err.Tag())
			}
			if err.Field() == "Message" {
				return f.Message(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
