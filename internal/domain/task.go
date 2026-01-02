package domain

import "time"


// TaskID is a strong type for task identifiers
type TaskID string

// Priority represents task importance
type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)


// Task represents a unit of work in the system
type Task struct {
	ID          TaskID     `json:"id"`
	Title       string     `json:"title"`
	Notes       string     `json:"notes,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
	Priority    Priority   `json:"priority"`
	CreatedAt   time.Time  `json:"created_at"`
	DueAt       *time.Time `json:"due_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Deleted     bool       `json:"deleted"`
}