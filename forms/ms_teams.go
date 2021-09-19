package forms

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

// MSTeamsForm ...
type MSTeamsForm struct{}

// CreateMSTeamsNotifyForm ...
type CreateMSTeamsNotifyForm struct {
	Title   string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Content string `form:"content" json:"content" binding:"required,min=3,max=1000"`
}

// CreateMSTeamsAlertForm ...
type CreateMSTeamsAlertForm struct {
	Title       string    `form:"title" json:"title" binding:"required,min=3,max=100"`
	Priority    int       `form:"priority" json:"priority" binding:"required,min=0,max=10"`
	MonitorName string    `form:"monitor_name" json:"monitor_name" binding:"required,min=3,max=100"`
	Description string    `form:"description" json:"description,omitempty" binding:"max=100"`
	CreateDate  time.Time `form:"create_date" json:"create_date,omitempty"`
}

// NotifyTitle ...
func (f MSTeamsForm) NotifyTitle(tag string, errMsg ...string) (message string) {
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

// NotifyContent ...
func (f MSTeamsForm) NotifyContent(tag string, errMsg ...string) (message string) {
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

// AlertTitle ...
func (f MSTeamsForm) AlertTitle(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter alert title"
		}
		return errMsg[0]
	case "min", "max":
		return "Alert title should be between 3 to 100 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// AlertPriority ...
func (f MSTeamsForm) AlertPriority(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter alert priority"
		}
		return errMsg[0]
	case "min", "max":
		return "Alert priority should be between 0 to 10"
	default:
		return "Something went wrong, please try again later"
	}
}

// AlertMonitorName ...
func (f MSTeamsForm) AlertMonitorName(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter alert monitor name"
		}
		return errMsg[0]
	case "min", "max":
		return "Alert source should be between 0 to 100"
	default:
		return "Something went wrong, please try again later"
	}
}

// AlertDescription ...
func (f MSTeamsForm) AlertDescription(tag string, _ ...string) (message string) {
	switch tag {
	case "min":
		return "Alert source should be between 0 to 100"
	default:
		return "Something went wrong, please try again later"
	}
}

// CreateNotify ...
func (f MSTeamsForm) CreateNotify(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Title" {
				return f.NotifyTitle(err.Tag())
			}
			if err.Field() == "Content" {
				return f.NotifyContent(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// CreateAlert ...
func (f MSTeamsForm) CreateAlert(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Title" {
				return f.AlertTitle(err.Tag())
			}
			if err.Field() == "Priority" {
				return f.AlertPriority(err.Tag())
			}
			if err.Field() == "Source" {
				return f.AlertMonitorName(err.Tag())
			}
			if err.Field() == "Description" {
				return f.AlertDescription(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
