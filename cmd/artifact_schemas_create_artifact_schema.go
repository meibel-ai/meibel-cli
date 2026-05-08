package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	artifactSchemasCreateArtifactSchemaData string
	artifactSchemasCreateArtifactSchemaInteractive bool
)

var artifactSchemasCreateArtifactSchemaCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Artifact Schema",
	Long:  `Create Artifact Schema`,
	Example: "meibel artifact-schemas create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateAgentArtifactRequest

		if artifactSchemasCreateArtifactSchemaData != "" {
			if err := json.Unmarshal([]byte(artifactSchemasCreateArtifactSchemaData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if artifactSchemasCreateArtifactSchemaInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("DisplayName").Description("Human-readable name of the artifact (letters, numbers, and spaces only). Converted to kebab-case internally.").Value(&body.DisplayName),
					huh.NewInput().Title("SchemaDef").Description("Schema definition").Value(&body.SchemaDef),
					huh.NewInput().Title("AdditionalProperties").Description(""),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.ArtifactSchemas.CreateArtifactSchema(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	artifactSchemasCmd.AddCommand(artifactSchemasCreateArtifactSchemaCmd)

	artifactSchemasCreateArtifactSchemaCmd.Flags().StringVarP(&artifactSchemasCreateArtifactSchemaData, "data", "d", "", "JSON data for the request body")
	artifactSchemasCreateArtifactSchemaCmd.Flags().BoolVarP(&artifactSchemasCreateArtifactSchemaInteractive, "interactive", "i", false, "use interactive form input")
}
