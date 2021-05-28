package gerror

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
)

const (
	maxStackDepth = 32
)

var (
	// goRootForFilter is used for stack filtering purpose.
	// Mainly for development environment.
	goRootForFilter = runtime.GOROOT()
)

func init() {
	if goRootForFilter != "" {
		goRootForFilter = strings.Replace(goRootForFilter, "\\", "/", -1)
	}
}

type stack []uintptr

func (s *stack) buffer(buffer *bytes.Buffer) {
	for _, pc := range *s {
		frame(pc).buffer(buffer)
	}
}

type frame uintptr

func (f frame) pc() uintptr { return uintptr(f) - 1 }

func (f frame) buffer(buffer *bytes.Buffer) {
	if fn := runtime.FuncForPC(f.pc()); fn != nil {
		file, line := fileLine(fn, f.pc())

		// Ignore GO ROOT paths.
		if l := len(goRootForFilter); l > 0 && len(file) >= l && file[0:l] == goRootForFilter {
			return
		}

		buffer.WriteString("\n")
		buffer.WriteString("\t")
		buffer.WriteString("at ")
		buffer.WriteString(name(fn))
		buffer.WriteString("\t")
		buffer.WriteString(file)
		buffer.WriteString(":")
		buffer.WriteString(strconv.Itoa(line))
	}
}

func fileLine(fn *runtime.Func, pc uintptr) (string, int) {
	if fn == nil {
		return "unknown", 0
	}
	return fn.FileLine(pc)
}

func name(fn *runtime.Func) string {
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

func callers() *stack {
	var pcs [maxStackDepth]uintptr
	var st stack = pcs[:runtime.Callers(3, pcs[:])]
	return &st
}
