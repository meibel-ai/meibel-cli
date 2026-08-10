package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
)

var documentsGetStructuredResultCmd = &cobra.Command{
	Use:   "get-structured-result <job-id>",
	Short: "Get structured parse result",
	Long:  `Download the fully structured parse result (the json format): pages, typed elements, tables, chart data, chart OCR text, and bounding boxes. The response schema (StructuredDocument) is defined by the parsing engine and hoisted into this spec by the OpenAPI generator.

Arguments:
  job-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel documents get-structured-result <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		sp := tui.StartSpinner("Get structured parse result")
		result, err := client.Documents.GetStructuredResult(ctx, jobId)
		sp.Stop()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsGetStructuredResultCmd)

}
