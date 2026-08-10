package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var documentsDownloadDeepTransformArtifactCmd = &cobra.Command{
	Use:   "download-deep-transform-artifact <job-id> <name>",
	Short: "Download a deep-transform artifact",
	Long:  `Download a named artifact (e.g. output.json) produced by a succeeded job. Ownership is verified against the customer header before any bytes are returned.

Arguments:
  job-id: required
  name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel documents download-deep-transform-artifact <job-id> <name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]
		name := args[1]

		sp := tui.StartSpinner("Download a deep-transform artifact")
		result, err := client.Documents.DownloadDeepTransformArtifact(ctx, jobId, name)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsDownloadDeepTransformArtifactCmd)

}
