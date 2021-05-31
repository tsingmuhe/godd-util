package gerror

import "fmt"

func New(msg string) error {
	return &runtimeException{
		stack: callers(),
		msg:   msg,
	}
}

func Newf(format string, args ...interface{}) error {
	return &runtimeException{
		stack: callers(),
		msg:   fmt.Sprintf(format, args...),
	}
}

func WithCause(err error, msg string) error {
	return &runtimeException{
		cause: err,
		stack: callers(),
		msg:   msg,
	}
}

func WithCausef(err error, format string, args ...interface{}) error {
	return &runtimeException{
		cause: err,
		stack: callers(),
		msg:   fmt.Sprintf(format, args...),
	}
}

func RootCause(err error) error {
	for err != nil {
		e, ok := err.(*runtimeException)
		if !ok {
			break
		}

		cause := e.cause
		if cause != nil {
			err = cause
		} else {
			return e
		}
	}

	return err
}
