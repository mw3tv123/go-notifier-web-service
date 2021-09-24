package forms

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

// MSTeamsForm ...
type MSTeamsForm struct{}

// CreateMSTeamsNotifyForm ...
type CreateMSTeamsNotifyForm struct {
	Title   string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Content string `form:"content" json:"content" binding:"required,min=3,max=200"`
}

// CreateMSTeamsAlertForm ...
type CreateMSTeamsAlertForm struct {
	Title       string    `form:"title" json:"title" binding:"required,min=3,max=100"`
	Priority    int       `form:"priority" json:"priority" binding:"required,min=0,max=10"`
	ServiceName string    `form:"service_name" json:"service_name" binding:"required,min=3,max=100"`
	Description string    `form:"description" json:"description,omitempty" binding:"max=100"`
	CreateDate  time.Time `form:"create_date" json:"create_date,omitempty"`
}

// evaluateErrorMessage evaluate error message by combine tag, field name and param
func (f MSTeamsForm) evaluateErrorMessage(tag, fieldName, param string, errMsg ...string) (message string) {
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
	default:
		return "Something went wrong, please try again later"
	}
}

// GetValidatedErrorMessage will return error message when form not pass validation
func (f MSTeamsForm) GetValidatedErrorMessage(err error) string {
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
