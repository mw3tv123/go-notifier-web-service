package forms

import (
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// DefaultValidator ...
type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

// ValidateStruct ...
func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine ...
func (v *DefaultValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

func (v *DefaultValidator) lazyInit() {
	v.once.Do(func() {

		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// Add any custom validations etc. here
		_ = v.validate.RegisterValidation("supportChannel", ValidateChannel)
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

// ValidateChannel will check if each channel request by client is supported by application or not.
func ValidateChannel(fl validator.FieldLevel) bool {
	channel := fl.Field().String()

	for _, c := range SupportedChannels {
		if c == channel {
			return true
		}
	}

	return false
}
