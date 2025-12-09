package domain

import (
	"time"
)

type TaskID string

type Priority string 
const(
	PriorityLow Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh Priority = "high"
)


type Task struct{
	ID 			TaskID  	`json:"id"`
	Title 		string	`json:"title"`
	Notes 		string	`json:"notes,omitempty"`
	Tags 		[]string	`json:"tags,omitempty"`
	Priority 	Priority	`json:"priority,omitempty"`
	CreatedAt 	time.Time	`json:"created_at"`
	DueAt 		*time.Time	`json:"due_at,omitempty"`
	CompletedAt *time.Time	`json:"completed_at,omitempty"`
	Deleted 	bool		`json:"deleted,omitempty"`
}