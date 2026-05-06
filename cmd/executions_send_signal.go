package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	executionsSendSignalData string
	executionsSendSignalInteractive bool
)

var executionsSendSignalCmd = &cobra.Command{
	Use:   "send-signal <blueprint-instance-id> <signal-name>",
	Short: "Send Signal",
	Long:  `Send Signal

Arguments:
  blueprint-instance-id: Unique identifier for the workflow instance
  signal-name: Name of the signal to send`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel blueprints executions send-signal <blueprint-instance-id> <signal-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]
		signalName := args[1]

		var body interface{}

		if executionsSendSignalData != "" {
			if err := json.Unmarshal([]byte(executionsSendSignalData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Blueprints.Executions.SendSignal(ctx, blueprintInstanceId, signalName, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	executionsCmd.AddCommand(executionsSendSignalCmd)

	executionsSendSignalCmd.Flags().StringVarP(&executionsSendSignalData, "data", "d", "", "JSON data for the request body")
	executionsSendSignalCmd.Flags().BoolVarP(&executionsSendSignalInteractive, "interactive", "i", false, "use interactive form input")
}
