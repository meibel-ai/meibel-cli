package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	confidenceScoringGetScoringJobsSummaryPrimary string
	confidenceScoringGetScoringJobsSummarySecondary string
)

var confidenceScoringGetScoringJobsSummaryCmd = &cobra.Command{
	Use:   "get-jobs-summary",
	Short: "Get Scoring Jobs Summary",
	Long:  `Get Scoring Jobs Summary`,
	Example: "meibel confidence-scoring get-jobs-summary --secondary=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.GetScoringJobsSummaryOptions{}
		if confidenceScoringGetScoringJobsSummarySecondary != "" {
			opts.Secondary = &confidenceScoringGetScoringJobsSummarySecondary
		}

		result, err := client.ConfidenceScoring.GetScoringJobsSummary(ctx, confidenceScoringGetScoringJobsSummaryPrimary, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	confidenceScoringCmd.AddCommand(confidenceScoringGetScoringJobsSummaryCmd)

	confidenceScoringGetScoringJobsSummaryCmd.Flags().StringVarP(&confidenceScoringGetScoringJobsSummaryPrimary, "primary", "", "", "The primary parameter")
	confidenceScoringGetScoringJobsSummaryCmd.MarkFlagRequired("primary")
	confidenceScoringGetScoringJobsSummaryCmd.Flags().StringVarP(&confidenceScoringGetScoringJobsSummarySecondary, "secondary", "", "", "The secondary parameter")
}
