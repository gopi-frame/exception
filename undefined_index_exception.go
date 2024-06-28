package exception

import "fmt"

type UndefinedIndexException struct {
	ArgumentException
}

func (e UndefinedIndexException) Unwrap() error {
	return e.ArgumentException
}

func NewUndefinedIndexException(index string) *UndefinedIndexException {
	exp := UndefinedIndexException{}
	exp.ArgumentException = ArgumentException{
		Throwable: New(fmt.Sprintf("Undefined index: \"%s\"", index)),
	}
	return &exp
}
