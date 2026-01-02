package ports

import "time"

// Clock provides the current time.
// Abstracted for testability and consistency.
type Clock interface {
	NowUTC() time.Time
}
