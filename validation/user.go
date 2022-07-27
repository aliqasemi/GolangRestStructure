package validation

import (
	"basical-app/models"
	"github.com/go-playground/validator"
	"time"
)

type (
	UserInput struct {
		FirstName       string `json:"first-name" form:"first-name" query:"first-name" param:"first-name" bson:"first-name,omitempty" validate:"required"`
		LastName        string `json:"last-name" form:"last-name" query:"last-name" param:"last-name" bson:"last-name,omitempty" validate:"required"`
		Age             int    `json:"age" form:"age" query:"age" param:"age" bson:"age,omitempty" validate:"gte=0,lte=130"`
		Password        string `json:"password" form:"password" query:"password" param:"password" bson:"password,omitempty" validate:"min=6,eqfield=ConfirmPassword"`
		ConfirmPassword string `json:"confirm-password" form:"confirm-password" query:"confirm-password" param:"confirm-password" bson:"confirm-password,omitempty"`
		PhoneNumber     string `json:"phone-number" form:"phone-number" query:"phone-number" param:"phone-number" bson:"phone-number,omitempty" validate:"min=10"`
		Email           string `json:"email" form:"email" query:"email" param:"email" bson:"email,omitempty" validate:"required,email"`
	}
	CustomValidator struct {
		validator *validator.Validate
	}
)

type UserValidation interface {
	validate(input UserInput) (bool, error)
	buildModel(input UserInput) (models.User, error)
}

func (input *UserInput) ValidateAndBuildModel() (models.User, error) {
	validatorInput := &CustomValidator{validator: validator.New()}
	validate, err := validatorInput.validate(input)
	if validate {
		return buildModel(input), nil
	} else {
		return models.User{}, err.(validator.ValidationErrors)
	}
}

func (validator *CustomValidator) validate(input *UserInput) (bool, error) {
	if err := validator.validator.Struct(input); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func buildModel(input *UserInput) models.User {
	return models.User{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Age:         input.Age,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		CreatedAt:   time.Now(),
	}
}
