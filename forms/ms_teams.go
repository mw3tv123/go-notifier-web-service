package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// NotifyForm ...
type NotifyForm struct{}

// CreateMSTeamNotifyForm ...
type CreateMSTeamNotifyForm struct {
	Title   string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Content string `form:"content" json:"content" binding:"required,min=3,max=1000"`
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

// Content ...
func (f NotifyForm) Content(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your notify content"
		}
		return errMsg[0]
	case "min", "max":
		return "Content should be between 3 to 1000 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// CreateNotify ...
func (f NotifyForm) CreateNotify(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Title" {
				return f.Title(err.Tag())
			}
			if err.Field() == "Content" {
				return f.Content(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
