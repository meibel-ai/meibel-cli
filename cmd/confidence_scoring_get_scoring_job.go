package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var confidenceScoringGetScoringJobCmd = &cobra.Command{
	Use:   "get-job <job-id>",
	Short: "Get Scoring Job",
	Long:  `Get Scoring Job

Arguments:
  job-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel confidence-scoring get-job <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		result, err := client.ConfidenceScoring.GetScoringJob(ctx, jobId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	confidenceScoringCmd.AddCommand(confidenceScoringGetScoringJobCmd)

}
