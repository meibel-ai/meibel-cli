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
	blueprintsCreateBlueprintTaskData string
	blueprintsCreateBlueprintTaskInteractive bool
)

var blueprintsCreateBlueprintTaskCmd = &cobra.Command{
	Use:   "create-task <blueprint-id>",
	Short: "Create Blueprint Task",
	Long:  `Create Blueprint Task

Arguments:
  blueprint-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints create-task <blueprint-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintId := args[0]

		var body sdk.AddBlueprintTaskRequest

		if blueprintsCreateBlueprintTaskData != "" {
			if err := json.Unmarshal([]byte(blueprintsCreateBlueprintTaskData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsCreateBlueprintTaskInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Name").Description("").Value(&body.Name),
					huh.NewInput().Title("InputSchema").Description("").Value(&body.InputSchema),
					huh.NewInput().Title("OutputSchema").Description("").Value(&body.OutputSchema),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Blueprints.CreateBlueprintTask(ctx, blueprintId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsCreateBlueprintTaskCmd)

	blueprintsCreateBlueprintTaskCmd.Flags().StringVarP(&blueprintsCreateBlueprintTaskData, "data", "d", "", "JSON data for the request body")
	blueprintsCreateBlueprintTaskCmd.Flags().BoolVarP(&blueprintsCreateBlueprintTaskInteractive, "interactive", "i", false, "use interactive form input")
}
