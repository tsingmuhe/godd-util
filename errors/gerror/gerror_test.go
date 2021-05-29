package gerror

import (
	"fmt"
	"testing"
)

type CodeError struct {
	*Exception
	Code int
}

func Test_CodeError(t *testing.T) {
	var e error = &CodeError{
		Exception: New("hello"),
		Code:      1,
	}
	fmt.Printf("%+v\n", e)
}
