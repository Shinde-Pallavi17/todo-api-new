package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var EmailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func RegisterCustomValidators(v *validator.Validate) {
	_ = v.RegisterValidation("customemail", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		return EmailRegex.MatchString(email)
	})
}
