package gerror

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

var (
	baseError = errors.New("test")
)

func Benchmark_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("test")
	}
}

func Benchmark_Newf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Newf("%s", "test")
	}
}

func Benchmark_Wrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithCause(baseError, "test")
	}
}

func Benchmark_Wrapf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithCausef(baseError, "%s", "test")
	}
}

func Benchmark_Stack(b *testing.B) {
	err := New("1")
	err = WithCause(err, "2")
	err = WithCause(err, "3")

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%+v", err)
	}
}

func Benchmark_Error(b *testing.B) {
	err := New("1")
	err = WithCause(err, "2")
	err = WithCause(err, "3")

	for i := 0; i < b.N; i++ {
		err.Error()
	}
}

func Test_New(t *testing.T) {
	err1 := New("test")
	fmt.Println(err1.Error())

	err2 := Newf("%v", "test")
	fmt.Println(err2.Error())
}

func Test_Wrap1(t *testing.T) {
	err := errors.New("1")
	err = WithCause(err, "2")
	err = WithCause(err, "3")

	fmt.Println(err.Error())
}

func Test_Wrap2(t *testing.T) {
	err := New("1")
	err = WithCause(err, "2")
	err = WithCause(err, "3")

	fmt.Println(err.Error())
}

func Test_Wrap3(t *testing.T) {
	err := New("1")
	err = WithCause(err, "")

	fmt.Println(err.Error())
}

func Test_Stack1(t *testing.T) {
	err := errors.New("1")
	fmt.Printf("%+v\n", err)
}

func Test_Stack2(t *testing.T) {
	err := errors.New("1")
	err = WithCause(err, "2")
	err = WithCause(err, "3")

	fmt.Printf("%+v\n", err)
}

func Test_Stack3(t *testing.T) {
	err := New("1")
	fmt.Printf("%+v\n", err)
}

func Test_Stack4(t *testing.T) {
	err := New("1")
	err = WithCause(err, "2")
	err = WithCause(err, "3")

	fmt.Printf("%+v\n", err)
}

func Test_Stack5(t *testing.T) {
	err := &runtimeException{
		msg: "hello",
	}

	fmt.Printf("%+v\n", err)
}

func Test_Json(t *testing.T) {
	err := WithCause(New("1"), "2")
	b, _ := json.Marshal(map[string]interface{}{"error": err})
	fmt.Println(string(b))
}

func Test_Null(t *testing.T) {
	var err error
	fmt.Printf("%+v\n", err)
	fmt.Printf("%v\n", err)
	fmt.Printf("%-v\n", err)

	b, _ := json.Marshal(map[string]interface{}{"error": err})
	fmt.Println(string(b))
}

func Test_Cause1(t *testing.T) {
	err := errors.New("1")
	err = WithCause(err, "2")
	err = WithCause(err, "3")

	e := RootCause(err)
	fmt.Printf("%+v\n", e)
}

func Test_Cause2(t *testing.T) {
	err := New("1")
	err = WithCause(err, "2")
	err = WithCause(err, "3")

	e := RootCause(err)
	fmt.Printf("%+v\n", e)
}
