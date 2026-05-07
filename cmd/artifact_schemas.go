package cmd

import (
	"github.com/spf13/cobra"
)

var artifactSchemasCmd = &cobra.Command{
	Use:   "artifact-schemas",
	Short: "Manage Artifact Schemas",
	Long:  `Commands for managing Artifact Schemas resources.`,
}

func init() {
	rootCmd.AddCommand(artifactSchemasCmd)
}
