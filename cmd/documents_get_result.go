package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	documentsGetResultFormat string
)

var documentsGetResultCmd = &cobra.Command{
	Use:   "get-result <job-id>",
	Short: "Get parsed document result",
	Long:  `Download the parsed result of a completed document parsing job.

Arguments:
  job-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel documents get-result <job-id> --format=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		opts := &sdk.DocumentsGetResultOptions{}
		if documentsGetResultFormat != "" {
			opts.Format = &documentsGetResultFormat
		}

		sp := tui.StartSpinner("Get parsed document result")
		result, err := client.Documents.GetResult(ctx, jobId, opts)
		sp.Stop()
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
	documentsCmd.AddCommand(documentsGetResultCmd)

	documentsGetResultCmd.Flags().StringVarP(&documentsGetResultFormat, "format", "", "markdown", "Result format: markdown, annotated, docling, json")
}
