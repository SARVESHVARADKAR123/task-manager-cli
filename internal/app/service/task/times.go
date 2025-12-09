package task

import (
    "time"
	"strings"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
    "github.com/SARVESHVARADKAR123/task-manager-cli/internal/timeutil"
)

func (s *TaskService) processDueDate(dueInput string) (*time.Time, error) {
    dueInput = strings.TrimSpace(dueInput)
    if dueInput == "" {
        return nil, nil
    }

    dueLocal, err := timeutil.ParseDateString(dueInput, s.loc)
    if err != nil {
        return nil, ErrInvalidInput
    }

    nowLocal := time.Now().In(s.loc)
    if dueLocal.Before(nowLocal) {
        return nil, ErrInvalidInput
    }

    utc := timeutil.ToUTC(dueLocal)
    return &utc, nil
}

func (s *TaskService) nowUTC() time.Time {
    return timeutil.ToUTC(time.Now().In(s.loc))
}

func (s *TaskService) toLocal(t *domain.Task) *domain.Task {
    if t == nil {
        return nil
    }

    created := t.CreatedAt.In(s.loc)

    var due *time.Time
    if t.DueAt != nil {
        v := t.DueAt.In(s.loc)
        due = &v
    }

    var completed *time.Time
    if t.CompletedAt != nil {
        v := t.CompletedAt.In(s.loc)
        completed = &v
    }

    return &domain.Task{
        ID:          t.ID,
        Title:       t.Title,
        Notes:       t.Notes,
        Tags:        t.Tags,
        Priority:    t.Priority,
        CreatedAt:   created,
        DueAt:       due,
        CompletedAt: completed,
        Deleted:     t.Deleted,
    }
}

func (s *TaskService) toLocalList(list []*domain.Task) []*domain.Task {
    out := make([]*domain.Task, 0, len(list))
    for _, t := range list {
        out = append(out, s.toLocal(t))
    }
    return out
}


func (s *TaskService) applyDueDateUpdate(t *domain.Task, dueInput *string) error {
    if dueInput == nil {
        return nil
    }
    dueUTC, err := s.processDueDate(*dueInput)
    if err != nil {
        return err
    }
    t.DueAt = dueUTC
    return nil
}
