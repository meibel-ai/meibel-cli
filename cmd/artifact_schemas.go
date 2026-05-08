package cmd

import (
	"github.com/spf13/cobra"
)

var artifactSchemasCmd = &cobra.Command{
	Use:   "artifact-schemas",
	Short: "Manage ArtifactSchemas",
	Long:  `Commands for managing ArtifactSchemas resources.`,
}

func init() {
	rootCmd.AddCommand(artifactSchemasCmd)
}
