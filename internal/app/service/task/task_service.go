package task

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/repo"
	"github.com/google/uuid"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/timeutil"
)


type TaskService struct {
    repo repo.TaskRepository
    loc  *time.Location
}

type EditTaskInput struct {
    Title    *string
    Notes    *string
    Tags     *[]string
    Priority *domain.Priority
    DueInput *string
}

func NewTaskService(r repo.TaskRepository, loc *time.Location) *TaskService {
    if r == nil {
        panic("nil repo passed to NewTaskService")
    }
    if loc == nil {
        loc = time.Local
    }
    return &TaskService{repo: r, loc: loc}
}


func (s *TaskService) AddTask(
    ctx context.Context,
    title string,
    tags []string,
    notes string,
    dueInput string,
    priority domain.Priority,
) (*domain.Task, error) {

    if err := validateTitle(title); err != nil {
        return nil, err
    }

    if err := validatePriority(priority); err != nil {
        return nil, err
    }

    cleanTags := sanitizeTags(tags)
    if len(cleanTags) > 10 {
        return nil, ErrInvalidInput
    }

    dueUTC, err := s.processDueDate(dueInput)
    if err != nil {
        return nil, err
    }

    createdUTC := s.nowUTC()

    task := &domain.Task{
        ID:        domain.TaskID(uuid.New().String()),
        Title:     strings.TrimSpace(title),
        Notes:     strings.TrimSpace(notes),
        Tags:      cleanTags,
        Priority:  priority,
        CreatedAt: createdUTC,
        DueAt:     dueUTC,
        Deleted:   false,
    }

    if err := s.repo.Create(ctx, task); err != nil {
        if errors.Is(err, repo.ErrConflict) {
            return nil, ErrAlreadyExists
        }
        return nil, err
    }

    return task, nil
}


func (s *TaskService) ListTasks(ctx context.Context, filter domain.TaskFilter) ([]*domain.Task, error) {
    tasks, err := s.repo.List(ctx, filter)
    if err != nil {
        return nil, err
    }
    return s.toLocalList(tasks), nil
}


func (s *TaskService) GetTask(ctx context.Context, id domain.TaskID) (*domain.Task, error) {
    t, err := s.repo.Get(ctx, id)
    if err != nil {
        if errors.Is(err, repo.ErrNotFound) {
            return nil, repo.ErrNotFound
        }
        return nil, err
    }

    return s.toLocal(t), nil
}

func (s *TaskService) EditTask(
    ctx context.Context,
    id domain.TaskID,
    in EditTaskInput,
) (*domain.Task, error) {

    // Load
    t, err := s.repo.Get(ctx, id)
    if err != nil {
        if errors.Is(err, repo.ErrNotFound) {
            return nil, repo.ErrNotFound
        }
        return nil, err
    }

    // Apply updates (each helper handles nil safely)
    if err := applyTitleUpdate(t, in.Title); err != nil {
        return nil, err
    }

    applyNotesUpdate(t, in.Notes)

    if err := applyTagsUpdate(t, in.Tags); err != nil {
        return nil, err
    }

    if err := applyPriorityUpdate(t, in.Priority); err != nil {
        return nil, err
    }

    if err := s.applyDueDateUpdate(t, in.DueInput); err != nil {
        return nil, err
    }

    // Persist
    if err := s.repo.Update(ctx, t); err != nil {
        return nil, err
    }

    return t, nil
}


func (s *TaskService) CompleteTask(ctx context.Context, id domain.TaskID) error {
    // 1. Load task
    t, err := s.repo.Get(ctx, id)
    if err != nil {
        if errors.Is(err, repo.ErrNotFound) {
            return repo.ErrNotFound
        }
        return err
    }

    // 2. Already completed?
    if t.CompletedAt != nil {
        return ErrAlreadyDone
    }

    // 3. Set CompletedAt in UTC
    nowLocal := time.Now().In(s.loc)
    completedUTC := timeutil.ToUTC(nowLocal)
    t.CompletedAt = &completedUTC

    // 4. Persist update
    if err := s.repo.Update(ctx, t); err != nil {
        return err
    }

    return nil
    
}

