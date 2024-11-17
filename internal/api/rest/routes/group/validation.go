package group

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"go-clean/internal/api/rest/utils/rest_error"
	"go-clean/internal/constants/validation_rules"
)

func checkError(err error) error {
	if err != nil {
		return rest_error.NewValidation(err.Error())
	}
	return nil
}

func validateAddDTO(dto *AddDTO) error {
	err := validation.ValidateStruct(dto,
		validation.Field(&dto.Name, validation.Required, validation.RuneLength(validation_rules.LenMinGroupName, validation_rules.LenMaxGroupName)),
	)
	return checkError(err)
}

func validatePatchDTO(dto *PatchDTO) error {
	err := validation.ValidateStruct(dto,
		validation.Field(&dto.Name, validation.RuneLength(validation_rules.LenMinGroupName, validation_rules.LenMaxGroupName)),
	)
	return checkError(err)
}
