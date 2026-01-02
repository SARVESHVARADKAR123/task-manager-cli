package task

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/app/service/task"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"	
)

func newDeleteCmd(svc *task.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <task-id>",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return svc.Delete(
				context.Background(),
				domain.TaskID(args[0]),
			)
		},
	}

	return cmd
}
