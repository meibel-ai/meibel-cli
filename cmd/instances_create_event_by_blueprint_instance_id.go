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
	instancesCreateEventByBlueprintInstanceIdData string
	instancesCreateEventByBlueprintInstanceIdInteractive bool
)

var instancesCreateEventByBlueprintInstanceIdCmd = &cobra.Command{
	Use:   "create-event-by-blueprint-id <blueprint-instance-id>",
	Short: "Create Event By Blueprint Instance Id",
	Long:  `Create Event By Blueprint Instance Id

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances create-event-by-blueprint-id <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body sdk.CustomEventRequest

		if instancesCreateEventByBlueprintInstanceIdData != "" {
			if err := json.Unmarshal([]byte(instancesCreateEventByBlueprintInstanceIdData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if instancesCreateEventByBlueprintInstanceIdInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("EventName").Description("Name of the custom event being logged.").Value(&body.EventName),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Blueprints.Instances.CreateEventByBlueprintInstanceId(ctx, blueprintInstanceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	instancesCmd.AddCommand(instancesCreateEventByBlueprintInstanceIdCmd)

	instancesCreateEventByBlueprintInstanceIdCmd.Flags().StringVarP(&instancesCreateEventByBlueprintInstanceIdData, "data", "d", "", "JSON data for the request body")
	instancesCreateEventByBlueprintInstanceIdCmd.Flags().BoolVarP(&instancesCreateEventByBlueprintInstanceIdInteractive, "interactive", "i", false, "use interactive form input")
}
