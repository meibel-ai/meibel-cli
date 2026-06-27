package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var confidenceScoringGetAgentSessionScoringSummaryCmd = &cobra.Command{
	Use:   "get-agent-session-summary <agent-name> <session-id>",
	Short: "Get agent session scoring summary",
	Long:  `Get an aggregated summary of confidence scores for a specific agent session.

Arguments:
  agent-name: Name of the agent.
  session-id: Agent session ID.`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel confidence-scoring get-agent-session-summary <agent-name> <session-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentName := args[0]
		sessionId := args[1]

		result, err := client.ConfidenceScoring.GetAgentSessionScoringSummary(ctx, agentName, sessionId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	confidenceScoringCmd.AddCommand(confidenceScoringGetAgentSessionScoringSummaryCmd)

}
