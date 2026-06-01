package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/tui"
)

var downloadsStreamProgressCmd = &cobra.Command{
	Use:   "stream-progress <job-id> <datasource-id>",
	Short: "Stream Download Progress",
	Long:  `Stream Download Progress

Arguments:
  job-id: required
  datasource-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources downloads stream-progress <job-id> <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]
		datasourceId := args[1]

		// Set up signal handling for graceful shutdown
		ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()

		stream, err := client.Datasources.Downloads.StreamProgress(ctx, jobId, datasourceId)
		if err != nil {
			return err
		}
		defer stream.Close()

		return tui.StreamEvents(ctx, stream)
	},
}

func init() {
	downloadsCmd.AddCommand(downloadsStreamProgressCmd)

}
