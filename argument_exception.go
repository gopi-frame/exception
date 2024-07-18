package exception

import (
	"fmt"
	"strings"

	ec "github.com/gopi-frame/contract/exception"
)

// ArgumentException argument exception
type ArgumentException struct {
	ec.Throwable
	name         string
	invalidValue any
}

func (e ArgumentException) Unwrap() error {
	return e.Throwable
}

// NewArgumentException new argument exception
func NewArgumentException(name string, invalidValue any, messages ...string) *ArgumentException {
	exp := ArgumentException{}
	exp.name = name
	exp.invalidValue = invalidValue
	message := fmt.Sprintf("Invalid argument (%s): %v", name, invalidValue)
	if len(messages) > 0 {
		message += fmt.Sprintf(": %s", strings.Join(messages, "\n"))
	}
	exp.Throwable = New(message)
	return &exp
}

// NewEmptyArgumentException argument empty
func NewEmptyArgumentException(name string) *ArgumentException {
	exp := new(ArgumentException)
	exp.name = name
	exp.Throwable = New(fmt.Sprintf("Invalid argument (%s): must not be empty or nil", name))
	return exp
}
