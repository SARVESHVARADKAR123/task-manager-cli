package	task

import (
	"strings"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)


func validateTitle(title string) error{
	if strings.TrimSpace(title)==""{
		return ErrInvalidInput
	}
	return nil
}


func validatePriority(p domain.Priority) error{
	
	switch p{
		case domain.PriorityLow, domain.PriorityMedium, domain.PriorityHigh:
			return nil
		default:
			return ErrInvalidInput
	}
}

func sanitizeTags(tags []string) []string {
	seen := make(map[string]struct{})
	out := make([]string, 0, len(tags))

	for _, t := range tags {
		trimmed := strings.TrimSpace(t)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		out = append(out, trimmed)
	}
	return out
}


