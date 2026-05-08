package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/tui"
)

var downloadsStreamDownloadProgressCmd = &cobra.Command{
	Use:   "stream-progress <datasource-id> <job-id>",
	Short: "Stream Download Progress",
	Long:  `Stream Download Progress

Arguments:
  datasource-id: required
  job-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources downloads stream-progress <datasource-id> <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		jobId := args[1]

		// Set up signal handling for graceful shutdown
		ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()

		stream, err := client.Datasources.Downloads.StreamDownloadProgress(ctx, datasourceId, jobId)
		if err != nil {
			return err
		}
		defer stream.Close()

		return tui.StreamEvents(ctx, stream)
	},
}

func init() {
	downloadsCmd.AddCommand(downloadsStreamDownloadProgressCmd)

}
