package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var artifactSchemasGetArtifactSchemaCmd = &cobra.Command{
	Use:   "get <artifact-id>",
	Short: "Get Artifact Schema",
	Long:  `Get Artifact Schema

Arguments:
  artifact-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel artifact-schemas get <artifact-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		artifactId := args[0]

		result, err := client.ArtifactSchemas.GetArtifactSchema(ctx, artifactId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	artifactSchemasCmd.AddCommand(artifactSchemasGetArtifactSchemaCmd)

}
