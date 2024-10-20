package response

type Info struct {
	Ok      bool   `json:"ok"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Bad struct {
	HttpCode int
	Info     *Info
}

func NewBadInfo(code string) *Info {
	return &Info{
		Ok:   false,
		Code: code,
	}
}

func NewBad(httpCode int, code string) *Bad {
	return &Bad{
		HttpCode: httpCode,
		Info:     NewBadInfo(code),
	}
}

func NewBadInfoValidation(message string) *Info {
	return &Info{
		Ok:      false,
		Code:    Validation,
		Message: message,
	}
}

func (e *Bad) Error() string {
	return e.Info.Message
}
