package cmd

import (
	"github.com/spf13/cobra"
)

var confidenceScoringCmd = &cobra.Command{
	Use:   "confidence-scoring",
	Short: "Manage Confidence Scoring",
	Long:  `Commands for managing Confidence Scoring resources.`,
}

func init() {
	rootCmd.AddCommand(confidenceScoringCmd)
}
