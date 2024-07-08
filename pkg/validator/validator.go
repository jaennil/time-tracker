package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var passportRegex = regexp.MustCompile(`^\d{4} \d{6}$`)

func NewValidator() (*validator.Validate, error) {
	validate := validator.New()
	err := validate.RegisterValidation("passport", validatePassport)
	if err != nil {
		return nil, err
	}

	return validate, nil
}

func validatePassport(fl validator.FieldLevel) bool {
	return passportRegex.MatchString(fl.Field().String())
}
