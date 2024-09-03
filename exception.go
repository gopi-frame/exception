package exception

import (
	"runtime"
	"strconv"
	"strings"

	ec "github.com/gopi-frame/contract/exception"
)

// Exception exception
type exception struct {
	cause      error
	message    string
	stackTrace string
}

// StackTrace returns stack trace
func (e *exception) StackTrace() string {
	return e.stackTrace
}

// Unwrap unwrap exception
func (e *exception) Unwrap() error {
	return e.cause
}

func (e *exception) Error() string {
	var str = new(strings.Builder)
	if strings.TrimSpace(e.message) != "" {
		str.WriteString(e.message)
	}
	if e.cause != nil {
		str.WriteByte(':')
		str.WriteByte(' ')
		str.WriteString(e.cause.Error())
	}
	return str.String()
}

func (e *exception) String() string {
	return e.Error()
}

func stack() string {
	var stackTrace = ""
	pcs := make([]uintptr, 32)
	n := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		str := &strings.Builder{}
		str.WriteString(frame.File)
		str.WriteByte('(')
		str.WriteString(strconv.Itoa(frame.Line))
		str.WriteByte(')')
		str.WriteByte(':')
		str.WriteString(frame.Function)
		str.WriteByte('\n')
		stackTrace += str.String()
		if !more {
			break
		}
	}
	return stackTrace
}

func New(messages ...string) ec.Throwable {
	e := exception{}
	e.message = strings.Join(messages, "\n")
	e.stackTrace = stack()
	return &e
}

func Wrap(err error) ec.Throwable {
	e := exception{}
	e.cause = err
	e.stackTrace = stack()
	return &e
}

// WithMessage new exception
func WithMessage(err error, message string) ec.Throwable {
	e := exception{}
	e.cause = err
	e.message = message
	e.stackTrace = stack()
	return &e
}
