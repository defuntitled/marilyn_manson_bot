package repository

import "fmt"

type OptimisticLockError struct {
	msg string
}

// NewGetTransactionStatusError creates a new GetTransactionStatusError instance with the provided message and error.
func NewOptimisticLockError(msg string) *OptimisticLockError {
	return &OptimisticLockError{
		msg: msg,
	}
}

func (e OptimisticLockError) Error() string {
	return fmt.Sprintf("error occured: %v", e.msg)
}
