package repo

import (
	"context"
	"errors"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)


var (
	ErrNotFound	= errors.New("task not found")
	ErrConflict	= errors.New("task conflict")
	ErrInternal = errors.New("internal repository error")
	ErrInvalidData = errors.New("invalid data")
)

type TaskRepository interface {
	Create(ctx context.Context, t *domain.Task) error
	Update(ctx context.Context, t *domain.Task) error
	Get(ctx context.Context, id domain.TaskID) (*domain.Task, error)
	Delete(ctx context.Context, id domain.TaskID) error
	List(ctx context.Context, filter domain.TaskFilter) ([]*domain.Task, error)
}


type FileRepoConfig struct {
	Path string
}

type FileData struct{
	Version int  			`json:"version"`
	Tasks 	[]*domain.Task 	`json:"tasks"`
}

// func NewFileRepo(config FileRepoConfig) (TaskRepository, error) {
// 	return &fileRepo{config: config}, nil
// }




