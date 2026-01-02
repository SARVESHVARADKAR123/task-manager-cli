package json

import "errors"

var (
	// ErrNotFound is returned when a task does not exist
	ErrNotFound = errors.New("task not found")

	// ErrConflict is returned when a task already exists
	ErrConflict = errors.New("task already exists")

	// ErrInvalidData is returned when stored JSON is corrupt or incompatible
	ErrInvalidData = errors.New("invalid repository data")
)
