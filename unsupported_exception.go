package exception

import ec "github.com/gopi-frame/contract/exception"

// UnsupportedException unsupported exception
type UnsupportedException struct {
	ec.Throwable
}

// NewUnsupportedException new unsupported exception
func NewUnsupportedException(message string) *UnsupportedException {
	return &UnsupportedException{
		Throwable: New("Unsupported operation: " + message),
	}
}
