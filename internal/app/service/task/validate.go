package task

import (
    "strings"
    "github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)

func validateTitle(title string) error {
    if strings.TrimSpace(title) == "" {
        return ErrInvalidInput
    }
    return nil
}

func validatePriority(p domain.Priority) error {
    switch p {
    case domain.PriorityLow, domain.PriorityMedium, domain.PriorityHigh:
        return nil
    default:
        return ErrInvalidInput
    }
}
