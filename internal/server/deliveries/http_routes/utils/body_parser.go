package utils

import (
	"easy-life-back-go/internal/server/deliveries/http_routes/constants"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validatorInstance = validator.New(validator.WithRequiredStructEnabled())

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ServerErrorDetails struct {
	Path    string
	Message string
}

type ClientErrorDetails struct {
	Code   string        `json:"code"`
	Errors []ErrorDetail `json:"errors"`
}

type ValidationErrorResponse struct {
	ServerError *ServerErrorDetails
	ClientError *ClientErrorDetails
}

type BodyParseError interface {
	Error() *ValidationErrorResponse
}

func newValidationErrorResponse(path, serverMessage, responseCode string, errors []ErrorDetail) *ValidationErrorResponse {
	return &ValidationErrorResponse{
		ServerError: &ServerErrorDetails{
			Path:    path,
			Message: serverMessage,
		},
		ClientError: &ClientErrorDetails{
			Code:   responseCode,
			Errors: errors,
		},
	}
}

func (e *ValidationErrorResponse) Error() *ValidationErrorResponse {
	return e
}

func BodyParser(r *http.Request, value any, shouldValidate bool) BodyParseError {
	if r.Body == nil {
		return newValidationErrorResponse(r.URL.Path, "request body is empty", http_errors_codes.RequestBodyIsEmpty, nil)
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(value); err != nil {
		return newValidationErrorResponse(r.URL.Path, "failed to decode JSON", http_errors_codes.InvalidJSON, nil)
	}

	if !shouldValidate {
		return nil
	}

	if err := validatorInstance.Struct(value); err != nil {
		var validationErrors []ErrorDetail

		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ErrorDetail{
				Field:   err.Field(),
				Message: fmt.Sprintf("%s validation failed for %s", err.Tag(), err.Param()),
			})
		}

		return newValidationErrorResponse(r.URL.Path, fmt.Sprintf("validation error: %s", err.Error()), http_errors_codes.Validation, validationErrors)
	}

	return nil
}
