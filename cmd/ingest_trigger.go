package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
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

		sp := tui.StartSpinner("Trigger Ingest")
		result, err := client.Datasources.Ingest.Trigger(ctx, datasourceId)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	ingestCmd.AddCommand(ingestTriggerCmd)

}
