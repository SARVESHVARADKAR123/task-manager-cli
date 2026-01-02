package ports

import(
	"context"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)

type TaskRepository interface{
	Create(ctx context.Context, task *domain.Task) error
	Update(ctx context.Context,task *domain.Task) error
	Get(ctx context.Context,id domain.TaskID) (*domain.Task,error)
	Delete(ctx context.Context,id domain.TaskID) error
	List(ctx context.Context,filter domain.TaskFilter) ([]*domain.Task,error)

}
	