package task

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/app/service/task"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)

func newCompleteCmd(svc *task.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "complete <task-id>",
		Short: "Mark a task as completed",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return svc.Complete(
				context.Background(),
				domain.TaskID(args[0]),
			)
		},
	}

	return cmd
}
