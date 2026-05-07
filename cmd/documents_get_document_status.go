package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var documentsGetDocumentStatusCmd = &cobra.Command{
	Use:   "get-status <job-id>",
	Short: "Get document parsing status",
	Long:  `Check the status of a document parsing job, including progress statistics.

Arguments:
  job-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel documents get-status <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		result, err := client.Documents.GetDocumentStatus(ctx, jobId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsGetDocumentStatusCmd)

}
