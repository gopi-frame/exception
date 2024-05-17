package exception

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/gopi-frame/contract/exception"
)

var _ exception.Throwable = new(Exception)

// Exception exception
type Exception struct {
	previous   exception.Throwable
	message    string
	stackTrace string
}

func (exception *Exception) Error() string {
	return exception.message
}

// SetMessage set error message
func (exception *Exception) SetMessage(message string) exception.Throwable {
	exception.message = message
	return exception
}

// SetStackTrace set stack trace
func (exception *Exception) SetStackTrace(stackTrace string) exception.Throwable {
	exception.stackTrace = stackTrace
	return exception
}

// StackTrace returns stack trace
func (exception *Exception) StackTrace() string {
	return exception.stackTrace
}

// SetPrevious set previous exception
func (exception *Exception) SetPrevious(e exception.Throwable) exception.Throwable {
	exception.previous = e
	return exception
}

// Previous returns previous exception
func (exception *Exception) Previous() exception.Throwable {
	return exception.previous
}

// NewException new Exception
func NewException(messages ...string) *Exception {
	var stackTrace = ""
	pcs := make([]uintptr, 16)
	n := runtime.Callers(0, pcs)
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
	return new(Exception).
		SetStackTrace(stackTrace).
		SetMessage(strings.Join(messages, "\n")).(*Exception)
}
