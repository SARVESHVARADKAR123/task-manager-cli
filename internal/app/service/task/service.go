package task

import(
	"context"
	"github.com/google/uuid"


	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/app/ports"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)

type Service struct {
	repo  ports.TaskRepository
	clock ports.Clock
}

func New(repo ports.TaskRepository, clock ports.Clock) *Service {
	if repo == nil || clock == nil {
		panic("nil dependency passed to task service")
	}
	return &Service{repo: repo, clock: clock}
}


func (s *Service) Add(
	ctx context.Context,
	title string,
	tags []string,
	notes string,
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

	now := s.clock.NowUTC()

	task := &domain.Task{
		ID:        domain.TaskID(uuid.NewString()),
		Title:     title,
		Notes:     notes,
		Tags:      cleanTags,
		Priority:  priority,
		CreatedAt: now,
		Deleted:   false,
	}

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}


//List tasks
func (s *Service) List(
	ctx context.Context,
	filter domain.TaskFilter,
) ([]*domain.Task, error) {

	return s.repo.List(ctx, filter)
}


//get tasks
func (s *Service) Get(
	ctx context.Context,
	id domain.TaskID,
) (*domain.Task, error) {

	task, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}
	return task, nil
}


//task completed 
func (s *Service) Complete(
	ctx context.Context,
	id domain.TaskID,
) error {

	task, err := s.repo.Get(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	if task.CompletedAt != nil {
		return ErrAlreadyDone
	}

	now := s.clock.NowUTC()
	task.CompletedAt = &now

	return s.repo.Update(ctx, task)
}


//soft delete

func (s *Service) Delete(
	ctx context.Context,
	id domain.TaskID,
) error {

	task, err := s.repo.Get(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	task.Deleted = true
	return s.repo.Update(ctx, task)
}

