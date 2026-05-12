package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var (
	agentsSessionsCreateData string
	agentsSessionsCreateInteractive bool
)

var agentsSessionsCreateCmd = &cobra.Command{
	Use:   "create <agent-id>",
	Short: "Create Session",
	Long:  `Create Session

Arguments:
  agent-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents agents-sessions create <agent-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentId := args[0]

		var body interface{}

		if agentsSessionsCreateData != "" {
			if err := json.Unmarshal([]byte(agentsSessionsCreateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Agents.Sessions.Create(ctx, agentId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsSessionsCmd.AddCommand(agentsSessionsCreateCmd)

	agentsSessionsCreateCmd.Flags().StringVarP(&agentsSessionsCreateData, "data", "d", "", "JSON data for the request body")
	agentsSessionsCreateCmd.Flags().BoolVarP(&agentsSessionsCreateInteractive, "interactive", "i", false, "use interactive form input")
}
