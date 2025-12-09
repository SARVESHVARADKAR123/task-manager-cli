package task

import (
    "strings"
    "github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)

func applyTitleUpdate(t *domain.Task, title *string) error {
    if title == nil {
        return nil
    }
    if err := validateTitle(*title); err != nil {
        return err
    }
    t.Title = strings.TrimSpace(*title)
    return nil
}

func applyNotesUpdate(t *domain.Task, notes *string) {
    if notes == nil {
        return
    }
    t.Notes = strings.TrimSpace(*notes)
}

func applyTagsUpdate(t *domain.Task, tags *[]string) error {
    if tags == nil {
        return nil
    }
    clean := sanitizeTags(*tags)
    if len(clean) > 10 {
        return ErrInvalidInput
    }
    t.Tags = clean
    return nil
}

func applyPriorityUpdate(t *domain.Task, priority *domain.Priority) error {
    if priority == nil {
        return nil
    }
    if err := validatePriority(*priority); err != nil {
        return err
    }
    t.Priority = *priority
    return nil
}
