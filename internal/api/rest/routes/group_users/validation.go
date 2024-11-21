package group_users

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go-clean/internal/api/rest/utils/rest_error"
)

func checkError(err error) error {
	if err != nil {
		return rest_error.NewValidation(err.Error())
	}
	return nil
}

func validateInviteUserDTO(dto *InviteUserDTO) error {
	err := validation.ValidateStruct(dto,
		validation.Field(&dto.UserID, validation.Required, is.UUIDv4),
	)
	return checkError(err)
}

func validateExcludeUserDTO(dto *ExcludeUserDTO) error {
	err := validation.ValidateStruct(dto,
		validation.Field(&dto.UserID, validation.Required, is.UUIDv4),
	)
	return checkError(err)
}
