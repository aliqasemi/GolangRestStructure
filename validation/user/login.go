package validation

import (
	"basical-app/models"
	"github.com/go-playground/validator"
	"time"
)

type (
	LoginInput struct {
		PhoneNumber string `json:"phone-number" form:"phone-number" query:"phone-number" param:"phone-number" bson:"phone-number,omitempty" validate:"min=10"`
		Password    string `json:"password" form:"password" query:"password" param:"password" bson:"password,omitempty" validate:"required"`
	}
	LoginValidator struct {
		validator *validator.Validate
	}
)

type LoginValidation interface {
	validate(input UserInput) (bool, error)
	buildModel(input UserInput) (models.User, error)
}

func (input *LoginInput) ValidateAndBuildModel() (models.User, error) {
	validatorInput := &LoginValidator{validator: validator.New()}
	validate, err := validatorInput.validate(input)
	if validate {
		return validatorInput.buildModel(input), nil
	} else {
		return models.User{}, err.(validator.ValidationErrors)
	}
}

func (validator *LoginValidator) validate(input *LoginInput) (bool, error) {
	if err := validator.validator.Struct(input); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (validator *LoginValidator) buildModel(input *LoginInput) models.User {
	return models.User{
		PhoneNumber: input.PhoneNumber,
		CreatedAt:   time.Now(),
	}
}
