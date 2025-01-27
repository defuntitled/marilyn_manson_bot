package server

import "fmt"

type FormatError struct {
	msg string
}

// NewGetTransactionStatusError creates a new GetTransactionStatusError instance with the provided message and error.
func NewOptimisticLockError(msg string) *FormatError {
	return &FormatError{
		msg: msg,
	}
}

func (e FormatError) Error() string {
	return fmt.Sprintf("error occured: %v", e.msg)
}
