package task

import "errors"

var (
	ErrInvalidInput  = errors.New("invalid input")
	ErrAlreadyExists = errors.New("task already exists")
	ErrNotFound      = errors.New("task not found")
	ErrAlreadyDone   = errors.New("task already completed")
)
