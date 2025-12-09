package repo

import "github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"

type TaskRepo interface {
	Add(task domain.Task) error
	List() ([]domain.Task, error)
	Delete(id string) error
} 