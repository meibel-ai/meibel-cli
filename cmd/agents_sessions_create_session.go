package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var (
	agentsSessionsCreateSessionData string
	agentsSessionsCreateSessionInteractive bool
)

var agentsSessionsCreateSessionCmd = &cobra.Command{
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

		if agentsSessionsCreateSessionData != "" {
			if err := json.Unmarshal([]byte(agentsSessionsCreateSessionData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Agents.Sessions.CreateSession(ctx, agentId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsSessionsCmd.AddCommand(agentsSessionsCreateSessionCmd)

	agentsSessionsCreateSessionCmd.Flags().StringVarP(&agentsSessionsCreateSessionData, "data", "d", "", "JSON data for the request body")
	agentsSessionsCreateSessionCmd.Flags().BoolVarP(&agentsSessionsCreateSessionInteractive, "interactive", "i", false, "use interactive form input")
}
