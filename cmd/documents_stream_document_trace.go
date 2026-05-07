package cmd

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var documentsStreamDocumentTraceCmd = &cobra.Command{
	Use:   "stream-trace <job-id>",
	Short: "Stream document parsing trace",
	Long:  `Subscribe to real-time parsing progress via Server-Sent Events.

Arguments:
  job-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel documents stream-trace <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		// Set up signal handling for graceful shutdown
		ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()

		stream, err := client.Documents.StreamDocumentTrace(ctx, jobId)
		if err != nil {
			return err
		}
		defer stream.Close()

		return tui.StreamEvents(ctx, stream)
	},
}

func init() {
	documentsCmd.AddCommand(documentsStreamDocumentTraceCmd)

}
