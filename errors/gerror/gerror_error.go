package gerror

import (
	"bytes"
	"fmt"
	"io"
)

const NIL_STR = "<nil>"

type Exception struct {
	msg   string
	cause error
	stack *stack
}

func (err *Exception) Error() string {
	if err == nil {
		return NIL_STR
	}

	errStr := err.msg
	if err.cause != nil {
		if errStr != "" {
			errStr += ": "
		}
		errStr += err.cause.Error()
	}
	return errStr
}

func (err *Exception) Msg() string {
	if err == nil {
		return NIL_STR
	}
	return err.msg
}

func (err *Exception) Cause() error {
	if err == nil {
		return nil
	}
	return err.cause
}

func (err *Exception) Stack() string {
	if err == nil {
		return NIL_STR
	}

	var buffer = bytes.NewBuffer(nil)
	var loop = err

	buffer.WriteString(loop.msg)
	if loop.stack != nil {
		loop.stack.buffer(buffer)
	}

	for loop.cause != nil {
		buffer.WriteString("\nCaused by: ")
		if e, ok := loop.cause.(*Exception); ok {
			loop = e
			buffer.WriteString(loop.msg)
			if loop.stack != nil {
				loop.stack.buffer(buffer)
			}
		} else {
			buffer.WriteString(loop.cause.Error())
			break
		}
	}

	return buffer.String()
}

// Format formats the frame according to the fmt.Formatter interface.
//
// %v, %s   : Print all the error string;
// %-v, %-s : Print current level error string;
// %+v, %+s : Print full stack error list;
func (err *Exception) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		switch {
		case s.Flag('-'):
			io.WriteString(s, err.Msg())
		case s.Flag('+'):
			io.WriteString(s, err.Stack())
		default:
			io.WriteString(s, err.Error())
		}
	}
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (err *Exception) MarshalJSON() ([]byte, error) {
	return []byte(`"` + err.Error() + `"`), nil
}
