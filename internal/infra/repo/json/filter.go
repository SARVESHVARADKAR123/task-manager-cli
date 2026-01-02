package json

import (
	"sort"
	"strings"
	"time"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)



func applyFilter(tasks []*domain.Task, f domain.TaskFilter) []*domain.Task {
	out := make([]*domain.Task, 0)

	for _, t := range tasks {
		if matches(t, f) {
			out = append(out, t)
		}
	}

	sortTasks(out, f)
	return paginate(out, f)
}





func checkOverdue(t *domain.Task, f domain.TaskFilter) bool {
	if f.Overdue != nil && *f.Overdue {
		if t.DueAt == nil || t.CompletedAt != nil {
			return false
		}
		if !t.DueAt.Before(time.Now().UTC()) {
			return false
		}
	}
	return true
}

func matches(t *domain.Task, f domain.TaskFilter) bool {
	if t.Deleted {
		return false
	}

	if f.CompletedOnly != nil {
		if *f.CompletedOnly && t.CompletedAt == nil {
			return false
		}
	}

	//check overdue
	if !checkOverdue(t, f) {
		return false
	}
	

	if f.TextContains != "" {
		q := strings.ToLower(f.TextContains)
		if !strings.Contains(strings.ToLower(t.Title), q) &&
			!strings.Contains(strings.ToLower(t.Notes), q) {
			return false
		}
	}

	return true
}


func sortTasks(tasks []*domain.Task, f domain.TaskFilter) {
	sort.SliceStable(tasks, func(i, j int) bool {
		a := tasks[i]
		b := tasks[j]

		var less bool

		switch f.SortBy {
		case "created_at":
			less = a.CreatedAt.Before(b.CreatedAt)

		case "due_at":
			if a.DueAt == nil {
				return false
			}
			if b.DueAt == nil {
				return true
			}
			less = a.DueAt.Before(*b.DueAt)

		case "priority":
			less = a.Priority < b.Priority

		default:
			// fallback: stable ID sort
			less = a.ID < b.ID
		}

		if f.Asc {
			return less
		}
		return !less
	})
}


func paginate(tasks []*domain.Task, f domain.TaskFilter) []*domain.Task {
	start := f.Offset
	if start < 0 {
		start = 0
	}

	if start >= len(tasks) {
		return []*domain.Task{}
	}

	end := len(tasks)
	if f.Limit > 0 && start+f.Limit < end {
		end = start + f.Limit
	}

	return tasks[start:end]
}
