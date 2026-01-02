package cli

import (
	"github.com/spf13/cobra"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/app/service/task"
	taskcmd "github.com/SARVESHVARADKAR123/task-manager-cli/internal/ui/cli/task"
)

func NewRoot(taskService *task.Service) *cobra.Command {
	root := &cobra.Command{
		Use:   "task-manager-cli",
		Short: "Task Manager CLI",
	}

	root.AddCommand(taskcmd.NewTaskCmd(taskService))
	return root
}
