package domain

import "time"

// TaskFilter represents criteria for querying tasks
type TaskFilter struct {
	ID            *TaskID
	Overdue       *bool
	DueBefore     *time.Time
	DueOn         *time.Time
	DueAfter      *time.Time
	TagsInclude   []string
	TextContains string
	CompletedOnly *bool
	Limit         int
	Offset        int
	SortBy        string
	Asc           bool
}