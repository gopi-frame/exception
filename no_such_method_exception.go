package exception

import (
	"fmt"
	"reflect"
)

// NoSuchMethodException no such method exception
type NoSuchMethodException struct {
	*Exception
	method       string
	receiverType reflect.Type
}

// NewNoSuchMethodException new no such method exception
func NewNoSuchMethodException(receiver any, method string) *NoSuchMethodException {
	exp := new(NoSuchMethodException)
	exp.method = method
	exp.receiverType = reflect.TypeOf(receiver)
	exp.Exception = NewException(fmt.Sprintf("NoSuchMethodException: no such method %s in %s", exp.method, exp.receiverType.Name()))
	return exp
}

// NewUnexportedMethodException new unexport method exception
func NewUnexportedMethodException(method string) *NoSuchMethodException {
	exp := new(NoSuchMethodException)
	exp.method = method
	exp.Exception = NewException(fmt.Sprintf("NoSuchMethodException: method %s is not exported", method))
	return exp
}
