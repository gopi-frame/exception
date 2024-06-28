package exception

import "github.com/gopi-frame/exception/contract"

// UnsupportedException unsupported exception
type UnsupportedException struct {
	contract.Throwable
}

// NewUnsupportedException new unsupported exception
func NewUnsupportedException(message string) *UnsupportedException {
	return &UnsupportedException{
		Throwable: New("Unsupported operation: " + message),
	}
}
