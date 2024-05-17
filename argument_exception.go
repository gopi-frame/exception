package exception

import (
	"fmt"
	"strings"
)

// ArgumentException argument exception
type ArgumentException struct {
	*Exception
	name         string
	invalidValue any
}

// NewArgumentException new argument exception
func NewArgumentException(name string, invalidValue any, messages ...string) *ArgumentException {
	exp := new(ArgumentException)
	exp.name = name
	exp.invalidValue = invalidValue
	message := fmt.Sprintf("Invalid argument (%s): %v", name, invalidValue)
	if len(messages) > 0 {
		message += fmt.Sprintf(": %s", strings.Join(messages, "\n"))
	}
	exp.Exception = NewException(message)
	return exp
}

// NewEmptyArgumentException argument empty
func NewEmptyArgumentException(name string) *ArgumentException {
	exp := new(ArgumentException)
	exp.name = name
	exp.Exception = NewException(fmt.Sprintf("Invalid argument (%s): must not be empty or nil", name))
	return exp
}
