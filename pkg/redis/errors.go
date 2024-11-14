package redis

const (
	codeNotFound = "not found"
)

// NotFoundError when key does not exist.
var NotFoundError = newError(codeNotFound)

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

func newError(msg string) error {
	return &Error{msg: msg}
}
