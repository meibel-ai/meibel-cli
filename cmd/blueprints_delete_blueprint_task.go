package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	blueprintsDeleteBlueprintTaskForce bool
)

var blueprintsDeleteBlueprintTaskCmd = &cobra.Command{
	Use:   "delete-task <blueprint-id> <task-id>",
	Short: "Delete Blueprint Task",
	Long:  `Delete Blueprint Task

Arguments:
  blueprint-id: required
  task-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel blueprints delete-task <blueprint-id> <task-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintId := args[0]
		taskId := args[1]

		if !blueprintsDeleteBlueprintTaskForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.Blueprints.DeleteBlueprintTask(ctx, blueprintId, taskId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsDeleteBlueprintTaskCmd)

	blueprintsDeleteBlueprintTaskCmd.Flags().BoolVarP(&blueprintsDeleteBlueprintTaskForce, "force", "f", false, "skip confirmation prompt")
}
