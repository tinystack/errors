package errors

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func New(text string) *Error {
	return newError(text, 3)
}

func Newf(format string, args ...interface{}) *Error {
	return newError(fmt.Sprintf(format, args...), 3)
}

func Wrap(err error, text string) *Error {
	e := newError(text, 3)
	e.Wrap(err)
	return e
}

func Wrapf(err error, format string, args ...interface{}) *Error {
	e := newError(fmt.Sprintf(format, args...), 3)
	e.Wrap(err)
	return e
}

func Is(err error) (e *Error, ok bool) {
	e, ok = err.(*Error)
	return
}

type Error struct {
	Message  string
	Cause    error
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
	Function string `json:"function,omitempty"`
}

var _ error = new(Error)

func newError(message string, skip int) *Error {
	e := &Error{
		Message: message,
	}
	e.fillCaller(skip)
	return e
}

func (e *Error) Error() string {
	var b bytes.Buffer
	if e.File != "" {
		b.WriteString(fmt.Sprintf("%s(%s:%d)", e.Function, e.File, e.Line))
	}
	b.WriteString(fmt.Sprintf(": %s", e.Message))
	if e.Cause != nil {
		b.WriteString(fmt.Sprintf(" | Caused: %s", e.Cause.Error()))
	}
	return b.String()
}

func (e *Error) Wrap(err error) {
	e.Cause = err
}

func (e *Error) fillCaller(skip int) {
	pc, fn, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	e.Line = line
	e.Function = funcName(pc)
	_, e.File = filepath.Split(fn)
}

func funcName(pc uintptr) string {
	if f := runtime.FuncForPC(pc); f != nil {
		name := f.Name()
		if index := strings.LastIndexByte(name, '/'); index != -1 {
			name = name[index+1:]
		}
		return name
	}
	return "??"
}
