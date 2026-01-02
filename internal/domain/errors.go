package domain

import "errors"

var (
	ErrTaskNotFound=errors.New("task not found")
	ErrTaskCompleted=errors.New("task already completed")
	ErrInvalidTask=errors.New("invalid task")
)

