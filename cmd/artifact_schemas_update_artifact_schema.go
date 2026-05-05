package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	artifactSchemasUpdateArtifactSchemaData string
	artifactSchemasUpdateArtifactSchemaInteractive bool
)

var artifactSchemasUpdateArtifactSchemaCmd = &cobra.Command{
	Use:   "update <artifact-id>",
	Short: "Update Artifact Schema",
	Long:  `Update Artifact Schema

Arguments:
  artifact-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel artifact-schemas update <artifact-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		artifactId := args[0]

		var body sdk.UpdateAgentArtifactRequest

		if artifactSchemasUpdateArtifactSchemaData != "" {
			if err := json.Unmarshal([]byte(artifactSchemasUpdateArtifactSchemaData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if artifactSchemasUpdateArtifactSchemaInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.ArtifactSchemas.UpdateArtifactSchema(ctx, artifactId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	artifactSchemasCmd.AddCommand(artifactSchemasUpdateArtifactSchemaCmd)

	artifactSchemasUpdateArtifactSchemaCmd.Flags().StringVarP(&artifactSchemasUpdateArtifactSchemaData, "data", "d", "", "JSON data for the request body")
	artifactSchemasUpdateArtifactSchemaCmd.Flags().BoolVarP(&artifactSchemasUpdateArtifactSchemaInteractive, "interactive", "i", false, "use interactive form input")
}
