package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var documentsGetDocumentResultCmd = &cobra.Command{
	Use:   "get-result <job-id>",
	Short: "Get parsed document result",
	Long:  `Download the parsed result of a completed document parsing job.

Arguments:
  job-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel documents get-result <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		result, err := client.Documents.GetDocumentResult(ctx, jobId)
		if err != nil {
			return err
		}

		if !output.PrintMarkdown(result, ".") {
			return output.Print(result)
		}
		return nil
	},
}

func init() {
	documentsCmd.AddCommand(documentsGetDocumentResultCmd)

}
