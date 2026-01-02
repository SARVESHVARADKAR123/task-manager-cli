package json

import (
	"context"
	"sync"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/app/ports"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)


type TaskRepo struct {
	path string
	mu   sync.RWMutex
}

func NewTaskRepo(path string) *TaskRepo {
	return &TaskRepo{path: path}
}

// Compile-time guarantee
var _ ports.TaskRepository = (*TaskRepo)(nil)


//Create 

func (r *TaskRepo) Create(ctx context.Context, task *domain.Task) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := r.load()
	if err != nil {
		return err
	}

	for _, t := range data.Tasks {
		if t.ID == task.ID {
			return ErrConflict
		}
	}

	data.Tasks = append(data.Tasks, task)
	return r.save(data)
}


//Get


func (r *TaskRepo) Get(ctx context.Context, id domain.TaskID) (*domain.Task, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	data, err := r.load()
	if err != nil {
		return nil, err
	}

	for _, t := range data.Tasks {
		if t.ID == id && !t.Deleted {
			cp := *t
			return &cp, nil
		}
	}

	return nil, ErrNotFound
}


//Update

func (r *TaskRepo) Update(ctx context.Context, task *domain.Task) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := r.load()
	if err != nil {
		return err
	}

	for i, t := range data.Tasks {
		if t.ID == task.ID && !t.Deleted {
			data.Tasks[i] = task
			return r.save(data)
		}
	}

	return ErrNotFound
}


//Delete (soft Delete)

func (r *TaskRepo) Delete(ctx context.Context, id domain.TaskID) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := r.load()
	if err != nil {
		return err
	}

	for i, t := range data.Tasks {
		if t.ID == id && !t.Deleted {
			data.Tasks[i].Deleted = true
			return r.save(data)
		}
	}

	return ErrNotFound
}


//List

func (r *TaskRepo) List(
	ctx context.Context,
	filter domain.TaskFilter,
) ([]*domain.Task, error) {

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	data, err := r.load()
	if err != nil {
		return nil, err
	}

	return applyFilter(data.Tasks, filter), nil
}
