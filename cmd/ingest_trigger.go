package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var ingestTriggerCmd = &cobra.Command{
	Use:   "trigger <datasource-id>",
	Short: "Trigger Ingest",
	Long:  `Trigger Ingest

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources ingest trigger <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.Datasources.Ingest.Trigger(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	ingestCmd.AddCommand(ingestTriggerCmd)

}
