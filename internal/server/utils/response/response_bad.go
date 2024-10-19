package response

type Bad struct {
	Ok      bool   `json:"ok"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewBad(code string) *Bad {
	return &Bad{
		Ok:   false,
		Code: code,
	}
}

func NewValidationError(message string) *Bad {
	return &Bad{
		Ok:      false,
		Code:    Validation,
		Message: message,
	}
}
