package validation

import (
	"easy-life-back-go/internal/constants/validation_rules"
	"easy-life-back-go/internal/server/routes/auth/views"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) SignIn(data *views.SignInData) error {
	err := validation.ValidateStruct(data,
		validation.Field(&data.Email, validation.Required, is.Email),
		validation.Field(&data.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
	)

	return err
}

func (v *Validator) Registration(data *views.RegistrationData) error {
	err := validation.ValidateStruct(data,
		validation.Field(&data.Email, validation.Required, is.Email),
	)

	return err
}

func (v *Validator) RegistrationSuccess(data *views.RegistrationSuccessData) error {
	err := validation.ValidateStruct(data,
		validation.Field(&data.Email, validation.Required, is.Email),
		validation.Field(&data.Name, validation.Required, validation.RuneLength(validation_rules.LenMinName, validation_rules.LenMaxName)),
		validation.Field(&data.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
		validation.Field(&data.Code, validation.Required, validation.RuneLength(validation_rules.LenRegistrationCode, validation_rules.LenRegistrationCode)),
	)

	return err
}
