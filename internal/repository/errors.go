package repository

import "fmt"

type OptimisticLockError struct {
	msg string
	err error
}

// NewGetTransactionStatusError creates a new GetTransactionStatusError instance with the provided message and error.
func NewOptimisticLockError(msg string, err error) *OptimisticLockError {
	return &OptimisticLockError{
		msg: msg,
		err: err,
	}
}

func (e OptimisticLockError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.err.Error())
}
