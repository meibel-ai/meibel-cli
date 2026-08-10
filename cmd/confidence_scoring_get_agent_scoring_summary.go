package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var confidenceScoringGetAgentScoringSummaryCmd = &cobra.Command{
	Use:   "get-agent-summary <agent-name>",
	Short: "Get agent scoring summary",
	Long:  `Get an aggregated summary of confidence scores for a specific agent.

Arguments:
  agent-name: Name of the agent to summarize.`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel confidence-scoring get-agent-summary <agent-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		agentName := args[0]

		sp := tui.StartSpinner("Get agent scoring summary")
		result, err := client.ConfidenceScoring.GetAgentScoringSummary(ctx, agentName)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	confidenceScoringCmd.AddCommand(confidenceScoringGetAgentScoringSummaryCmd)

}
