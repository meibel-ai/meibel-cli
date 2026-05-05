package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var contentTriggerIngestCmd = &cobra.Command{
	Use:   "trigger-ingest <datasource-id>",
	Short: "Trigger Ingest",
	Long:  `Trigger Ingest

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel content trigger-ingest <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.Content.TriggerIngest(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentTriggerIngestCmd)

}
