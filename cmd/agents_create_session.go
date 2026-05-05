package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var (
	agentsCreateSessionData string
	agentsCreateSessionInteractive bool
)

var agentsCreateSessionCmd = &cobra.Command{
	Use:   "create-session <agent-id>",
	Short: "Create Session",
	Long:  `Create Session

Arguments:
  agent-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents create-session <agent-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentId := args[0]

		var body interface{}

		if agentsCreateSessionData != "" {
			if err := json.Unmarshal([]byte(agentsCreateSessionData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Agents.CreateSession(ctx, agentId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	agentsCmd.AddCommand(agentsCreateSessionCmd)

	agentsCreateSessionCmd.Flags().StringVarP(&agentsCreateSessionData, "data", "d", "", "JSON data for the request body")
	agentsCreateSessionCmd.Flags().BoolVarP(&agentsCreateSessionInteractive, "interactive", "i", false, "use interactive form input")
}
