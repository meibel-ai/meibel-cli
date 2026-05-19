package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	confidenceScoringGetScoringJobsSummaryPrimary string
	confidenceScoringGetScoringJobsSummarySecondary string
)

var confidenceScoringGetScoringJobsSummaryCmd = &cobra.Command{
	Use:   "get-jobs-summary",
	Short: "Get scoring summary",
	Long:  `Get an aggregated summary of confidence scores. Requires a primary filter; an optional secondary filter narrows results further. Filters use the format "field:value", where field is any identity context field name.`,
	Example: "meibel confidence-scoring get-jobs-summary --secondary=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ConfidenceScoringGetScoringJobsSummaryOptions{}
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

	confidenceScoringGetScoringJobsSummaryCmd.Flags().StringVarP(&confidenceScoringGetScoringJobsSummaryPrimary, "primary", "", "", "Primary filter in \"field:value\" format, where field is an identity context field name (e.g. \"agent_name:my-agent\" or \"agent_session_id:sess_abc123\").")
	confidenceScoringGetScoringJobsSummaryCmd.MarkFlagRequired("primary")
	confidenceScoringGetScoringJobsSummaryCmd.Flags().StringVarP(&confidenceScoringGetScoringJobsSummarySecondary, "secondary", "", "", "Optional secondary filter in the same \"field:value\" format to further narrow results (e.g. \"agent_version:1.2.0\").")
}
