package gerror

import (
	"bytes"
	"fmt"
	"io"
)

type runtimeException struct {
	msg   string
	cause error
	stack *stack
}

func (e *runtimeException) stackString() string {
	var buffer = bytes.NewBuffer(nil)
	var loop = e

	buffer.WriteString(loop.msg)
	if loop.stack != nil {
		loop.stack.buffer(buffer)
	}

	for loop.cause != nil {
		buffer.WriteString("\nCaused by: ")
		if err, ok := loop.cause.(*runtimeException); ok {
			loop = err
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

func (e *runtimeException) Error() string {
	return e.msg
}

func (e *runtimeException) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'v':
		switch {
		case s.Flag('+'):
			io.WriteString(s, e.stackString())
		default:
			io.WriteString(s, e.Error())
		}
	}
}

func (e *runtimeException) MarshalJSON() ([]byte, error) {
	return []byte(`"` + e.Error() + `"`), nil
}
