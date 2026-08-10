package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	datasourcesCreateName string
	datasourcesCreateDescription string
	datasourcesCreateConnector string
	datasourcesCreateMetadataConfig string
)

var datasourcesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Datasource",
	Long:  `Create Datasource`,
	Example: "meibel datasources create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := sdk.DatasourcesCreateOptions{}
		opts.Name = datasourcesCreateName
		if datasourcesCreateDescription != "" {
			opts.Description = &datasourcesCreateDescription
		}
		if datasourcesCreateConnector != "" {
			opts.Connector = datasourcesCreateConnector
		}
		if datasourcesCreateMetadataConfig != "" {
			opts.MetadataConfig = datasourcesCreateMetadataConfig
		}

		result, err := client.Datasources.Create(ctx, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesCreateCmd)

	datasourcesCreateCmd.Flags().StringVar(&datasourcesCreateName, "name", "", "Human-readable datasource name")
	datasourcesCreateCmd.MarkFlagRequired("name")
	datasourcesCreateCmd.Flags().StringVar(&datasourcesCreateDescription, "description", "", "What this datasource contains")
	datasourcesCreateCmd.Flags().StringVar(&datasourcesCreateConnector, "connector", "", "Connection configuration — omit for file-upload datasources")
	datasourcesCreateCmd.Flags().StringVar(&datasourcesCreateMetadataConfig, "metadata-config", "", "Optional metadata extraction config to apply after creation")
}
