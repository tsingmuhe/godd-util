package gerror

import "fmt"

// New creates and returns an error which is formatted from given text.
func New(msg string) error {
	return &Error{
		stack: callers(),
		msg:   msg,
	}
}

// Newf returns an error that formats as the given format and args.
func Newf(format string, args ...interface{}) error {
	return &Error{
		stack: callers(),
		msg:   fmt.Sprintf(format, args...),
	}
}

// Wrap wraps error with text.
// It returns nil if given err is nil.
func Wrap(err error, msg string) error {
	return &Error{
		cause: err,
		stack: callers(),
		msg:   msg,
	}
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// It returns nil if given <err> is nil.
func Wrapf(err error, format string, args ...interface{}) error {
	return &Error{
		cause: err,
		stack: callers(),
		msg:   fmt.Sprintf(format, args...),
	}
}

func RootCause(err error) error {
	for err != nil {
		e, ok := err.(*Error)
		if !ok {
			break
		}

		cause := e.Cause()
		if cause != nil {
			err = cause
		} else {
			return e
		}
	}

	return err
}
