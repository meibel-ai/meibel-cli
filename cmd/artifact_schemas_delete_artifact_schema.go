package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	artifactSchemasDeleteArtifactSchemaForce bool
)

var artifactSchemasDeleteArtifactSchemaCmd = &cobra.Command{
	Use:   "delete <artifact-id>",
	Short: "Delete Artifact Schema",
	Long:  `Delete Artifact Schema

Arguments:
  artifact-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel artifact-schemas delete <artifact-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		artifactId := args[0]

		if !artifactSchemasDeleteArtifactSchemaForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		err := client.ArtifactSchemas.DeleteArtifactSchema(ctx, artifactId)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	artifactSchemasCmd.AddCommand(artifactSchemasDeleteArtifactSchemaCmd)

	artifactSchemasDeleteArtifactSchemaCmd.Flags().BoolVarP(&artifactSchemasDeleteArtifactSchemaForce, "force", "f", false, "skip confirmation prompt")
}
