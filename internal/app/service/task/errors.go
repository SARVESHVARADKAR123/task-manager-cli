package task


import "errors"

var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrAlreadyExists  = errors.New("task already exists")
	ErrAlreadyDone    = errors.New("task already completed")
	// ErrNotFound       = errors.New("task not found")
)
