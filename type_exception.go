package exception

import ec "github.com/gopi-frame/contract/exception"

// TypeException type exception
type TypeException struct {
	ec.Throwable
}

func (e TypeException) Unwrap() error {
	return e.Throwable
}

// NewTypeException new type exception
func NewTypeException(message string) *TypeException {
	exp := TypeException{}
	exp.Throwable = New("TypeException: " + message)
	return &exp
}
