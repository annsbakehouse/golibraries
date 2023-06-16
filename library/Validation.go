package library

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorMessage struct {
	Field string `json:"field"`
	Msg   string `json:"name"`
}

func ValidationErrorToText(field string, tag string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", field, param)
	case "min":
		return fmt.Sprintf("%s must be longer than %s", field, param)
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", field, param)
	case "smart_al_num_space":
		return fmt.Sprintf("%s must should contains Alphanumeric or space only", field)
	case "smart_num":
		return fmt.Sprintf("%s must be Number", field)
	case "smart_trim_space":
		return fmt.Sprintf("%s cannot started and terminated with space", field)
	}
	return fmt.Sprintf("%s is not valid", field)
}

var Validation_alphanumericspace validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	lnvalue := len(value)
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	clear := nonAlphanumericRegex.ReplaceAllString(value, "")
	if lnvalue != len(clear) {
		return false
	}
	return true
}
var Validation_number validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	lnvalue := len(value)
	var nonAlphanumericRegex = regexp.MustCompile(`[^0-9]+`)
	clear := nonAlphanumericRegex.ReplaceAllString(value, "")
	if lnvalue != len(clear) {
		return false
	}
	return true
}
var Validation_TrimSpace validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	lnvalue := len(value)
	clear := strings.TrimSpace(value)
	if lnvalue != len(clear) {
		return false
	}
	return true
}
