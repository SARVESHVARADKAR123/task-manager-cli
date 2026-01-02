package task

import (
	"github.com/spf13/cobra"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/app/service/task"
)

func NewTaskCmd(svc *task.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "Manage tasks",
	}

	cmd.AddCommand(
		newAddCmd(svc),
		newListCmd(svc),
		newCompleteCmd(svc),
		newDeleteCmd(svc),
	)

	return cmd
}
