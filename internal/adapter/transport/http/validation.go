package http

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"sync"
)

var customValidators = map[string]func(s validator.FieldLevel) bool{
	"email": func(s validator.FieldLevel) bool {
		id := s.Field().String()
		regex := regexp.MustCompile("^[a-f0-9]{24}$")
		return regex.MatchString(id)
	},
	"password": func(s validator.FieldLevel) bool {
		password := s.Field().String()

		uppercaseRegex := regexp.MustCompile(".*[A-Z].*")
		if !uppercaseRegex.MatchString(password) {
			return false
		}

		lowercaseRegex := regexp.MustCompile(".*[a-z].*")
		if !lowercaseRegex.MatchString(password) {
			return false
		}

		numericRegex := regexp.MustCompile(".*\\d.*")
		if !numericRegex.MatchString(password) {
			return false
		}

		specialRegex := regexp.MustCompile(".*[@*#$%^&+=!].*")
		if !specialRegex.MatchString(password) {
			return false
		}

		lengthRegex := regexp.MustCompile(".{8,20}")
		if !lengthRegex.MatchString(password) {
			return false
		}

		return true
	},
	"name": func(s validator.FieldLevel) bool {
		name := s.Field().String()
		regex := regexp.MustCompile("^[a-zA-Z ]+$")
		return regex.MatchString(name)
	},
	"surname": func(s validator.FieldLevel) bool {
		surname := s.Field().String()
		regex := regexp.MustCompile("^[a-zA-Z ]+$")
		return regex.MatchString(surname)
	},
	"phone": func(s validator.FieldLevel) bool {
		phone := s.Field().String()
		regex := regexp.MustCompile("^[0-9]{10}$")
		return regex.MatchString(phone)
	},
}

var validations *validator.Validate
var once sync.Once

func CreateNewValidator() *validator.Validate {
	once.Do(func() {
		validations = validator.New()
	})

	for key, value := range customValidators {
		validations.RegisterValidation(key, value)
	}

	return validations
}

func ValidateRequestByStruct[T any](s T) []*ValidationMessage {
	validate := CreateNewValidator()
	var errors []*ValidationMessage
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationMessage
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Message = err.Error()
			errors = append(errors, &element)
		}
	}
	return errors
}
