package task

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/app/service/task"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)

func newListCmd(svc *task.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := svc.List(
				context.Background(),
				domain.TaskFilter{},
			)
			if err != nil {
				return err
			}

			for _, t := range tasks {
				fmt.Printf("%s  %s\n", t.ID, t.Title)
			}
			return nil
		},
	}

	return cmd
}
