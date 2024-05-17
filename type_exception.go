package exception

// TypeException type exception
type TypeException struct {
	*Exception
}

// NewTypeException new type exception
func NewTypeException(message string) *TypeException {
	exp := new(TypeException)
	exp.Exception = NewException("TypeException: " + message)
	return exp
}
