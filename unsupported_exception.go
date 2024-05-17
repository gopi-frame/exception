package exception

// UnsupportedException unsupported exception
type UnsupportedException struct {
	*Exception
}

// NewUnsupportedException new unsupported exception
func NewUnsupportedException(message string) *UnsupportedException {
	return &UnsupportedException{
		Exception: NewException("Unsupported operation: " + message),
	}
}
