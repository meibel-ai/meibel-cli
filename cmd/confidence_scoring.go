package cmd

import (
	"github.com/spf13/cobra"
)

var confidenceScoringCmd = &cobra.Command{
	Use:   "confidence-scoring",
	Short: "Manage ConfidenceScoring",
	Long:  `Commands for managing ConfidenceScoring resources.`,
}

func init() {
	rootCmd.AddCommand(confidenceScoringCmd)
}
