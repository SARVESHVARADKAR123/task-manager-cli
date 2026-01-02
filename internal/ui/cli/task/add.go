package task

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/app/service/task"
	"github.com/SARVESHVARADKAR123/task-manager-cli/internal/domain"
)

func newAddCmd(svc *task.Service) *cobra.Command {
	var title string
	var notes string
	var priority string
	var tags []string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new task",
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := svc.Add(
				context.Background(),
				title,
				tags,
				notes,
				domain.Priority(priority),
			)
			if err != nil {
				return err
			}

			fmt.Println("Task created:", t.ID)
			return nil
		},
	}

	cmd.Flags().StringVar(&title, "title", "", "Task title (required)")
	cmd.Flags().StringVar(&notes, "notes", "", "Task notes")
	cmd.Flags().StringVar(&priority, "priority", "medium", "Priority: low|medium|high")
	cmd.Flags().StringSliceVar(&tags, "tags", nil, "Task tags")

	_ = cmd.MarkFlagRequired("title")
	return cmd
}
