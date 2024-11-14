package client_error

type ClientError struct {
	code string
}

func (err *ClientError) Error() string {
	return err.code
}

func (err *ClientError) Code() string {
	return err.code
}

func New(code string) *ClientError {
	return &ClientError{code: code}
}
