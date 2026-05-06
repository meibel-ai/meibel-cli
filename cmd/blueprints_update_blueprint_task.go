package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	blueprintsUpdateBlueprintTaskData string
	blueprintsUpdateBlueprintTaskInteractive bool
)

var blueprintsUpdateBlueprintTaskCmd = &cobra.Command{
	Use:   "update-task <blueprint-id> <task-id>",
	Short: "Update Blueprint Task",
	Long:  `Update Blueprint Task

Arguments:
  blueprint-id: required
  task-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel blueprints update-task <blueprint-id> <task-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintId := args[0]
		taskId := args[1]

		var body sdk.UpdateBlueprintTaskRequest

		if blueprintsUpdateBlueprintTaskData != "" {
			if err := json.Unmarshal([]byte(blueprintsUpdateBlueprintTaskData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsUpdateBlueprintTaskInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Blueprints.UpdateBlueprintTask(ctx, blueprintId, taskId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsUpdateBlueprintTaskCmd)

	blueprintsUpdateBlueprintTaskCmd.Flags().StringVarP(&blueprintsUpdateBlueprintTaskData, "data", "d", "", "JSON data for the request body")
	blueprintsUpdateBlueprintTaskCmd.Flags().BoolVarP(&blueprintsUpdateBlueprintTaskInteractive, "interactive", "i", false, "use interactive form input")
}
