package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	artifactSchemasCreateDisplayName string
	artifactSchemasCreateType string
	artifactSchemasCreateDescription string
	artifactSchemasCreateRequired string
	artifactSchemasCreateSchema string
	artifactSchemasCreateMaxSizeBytes string
	artifactSchemasCreateStorageStrategy string
	artifactSchemasCreateAdditionalProperties string
)

var artifactSchemasCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Artifact Schema",
	Long:  `Create Artifact Schema`,
	Example: "meibel artifact-schemas create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := sdk.ArtifactSchemasCreateOptions{}
		opts.DisplayName = artifactSchemasCreateDisplayName
		if artifactSchemasCreateType != "" {
			v := sdk.ArtifactType(artifactSchemasCreateType)
			opts.Type = &v
		}
		if artifactSchemasCreateDescription != "" {
			opts.Description = artifactSchemasCreateDescription
		}
		if artifactSchemasCreateRequired != "" {
			opts.Required = artifactSchemasCreateRequired
		}
		opts.Schema = artifactSchemasCreateSchema
		if artifactSchemasCreateMaxSizeBytes != "" {
			opts.MaxSizeBytes = artifactSchemasCreateMaxSizeBytes
		}
		if artifactSchemasCreateStorageStrategy != "" {
			opts.StorageStrategy = artifactSchemasCreateStorageStrategy
		}
		if artifactSchemasCreateAdditionalProperties != "" {
			var v map[string]interface{}
			if err := json.Unmarshal([]byte(artifactSchemasCreateAdditionalProperties), &v); err != nil {
				return fmt.Errorf("invalid JSON for --additional-properties: %w", err)
			}
			opts.AdditionalProperties = v
		}

		result, err := client.ArtifactSchemas.Create(ctx, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	artifactSchemasCmd.AddCommand(artifactSchemasCreateCmd)

	artifactSchemasCreateCmd.Flags().StringVar(&artifactSchemasCreateDisplayName, "display-name", "", "Human-readable name of the artifact (letters, numbers, and spaces only). Converted to kebab-case internally.")
	artifactSchemasCreateCmd.MarkFlagRequired("display-name")
	artifactSchemasCreateCmd.Flags().StringVar(&artifactSchemasCreateType, "type", "", "Artifact type (json, markdown, csv, yaml, text, html, pdf)")
	artifactSchemasCreateCmd.Flags().StringVar(&artifactSchemasCreateDescription, "description", "", "Description of the artifact")
	artifactSchemasCreateCmd.Flags().StringVar(&artifactSchemasCreateRequired, "required", "", "Whether agent must produce this artifact")
	artifactSchemasCreateCmd.Flags().StringVar(&artifactSchemasCreateSchema, "schema", "", "Schema definition")
	artifactSchemasCreateCmd.MarkFlagRequired("schema")
	artifactSchemasCreateCmd.Flags().StringVar(&artifactSchemasCreateMaxSizeBytes, "max-size-bytes", "", "Maximum artifact size in bytes")
	artifactSchemasCreateCmd.Flags().StringVar(&artifactSchemasCreateStorageStrategy, "storage-strategy", "", "Storage strategy (inline, gcs, auto)")
	artifactSchemasCreateCmd.Flags().StringVar(&artifactSchemasCreateAdditionalProperties, "additional-properties", "", "AdditionalProperties")
}
