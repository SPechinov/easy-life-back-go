package auth

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"go-clean/internal/api/rest/utils/rest_error"
	"go-clean/internal/constants/validation_rules"
	"go-clean/pkg/validate"
)

func checkError(err error) error {
	if err != nil {
		return rest_error.NewValidation(err.Error())
	}
	return nil
}

func validateLoginDTO(dto *LoginDTO) error {
	err := validation.ValidateStruct(dto,
		validation.Field(&dto.Email, validation.When(dto.Phone == "", validation.Required, validation.By(validate.Email))),
		validation.Field(&dto.Phone, validation.When(dto.Email == "", validation.Required, validation.By(validate.Phone))),
		validation.Field(&dto.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
	)
	return checkError(err)
}

func validateRegistrationDTO(dto *RegistrationDTO) error {
	err := validation.ValidateStruct(dto,
		validation.Field(&dto.Email, validation.When(dto.Phone == "", validation.Required, validation.By(validate.Email))),
		validation.Field(&dto.Phone, validation.When(dto.Email == "", validation.Required, validation.By(validate.Phone))),
	)
	return checkError(err)
}

func validateRegistrationConfirmDTO(dto *RegistrationConfirmDTO) error {
	err := validation.ValidateStruct(dto,
		validation.Field(&dto.Email, validation.When(dto.Phone == "", validation.Required, validation.By(validate.Email))),
		validation.Field(&dto.Phone, validation.When(dto.Email == "", validation.Required, validation.By(validate.Phone))),
		validation.Field(&dto.FirstName, validation.Required, validation.RuneLength(validation_rules.LenMinFirstName, validation_rules.LenMaxFirstName)),
		validation.Field(&dto.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
		validation.Field(&dto.Code, validation.Required, validation.RuneLength(validation_rules.LenRegistrationCode, validation_rules.LenRegistrationCode)),
	)
	return checkError(err)
}

func validateForgotPasswordDTO(dto *ForgotPasswordDTO) error {
	err := validation.ValidateStruct(dto,
		validation.Field(&dto.Email, validation.When(dto.Phone == "", validation.Required, validation.By(validate.Email))),
		validation.Field(&dto.Phone, validation.When(dto.Email == "", validation.Required, validation.By(validate.Phone))),
	)
	return checkError(err)
}

func validateForgotPasswordConfirmDTO(dto *ForgotPasswordConfirmDTO) error {
	err := validation.ValidateStruct(dto,
		validation.Field(&dto.Email, validation.When(dto.Phone == "", validation.Required, validation.By(validate.Email))),
		validation.Field(&dto.Phone, validation.When(dto.Email == "", validation.Required, validation.By(validate.Phone))),
		validation.Field(&dto.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
		validation.Field(&dto.Code, validation.Required, validation.RuneLength(validation_rules.LenRegistrationCode, validation_rules.LenRegistrationCode)),
	)
	return checkError(err)
}