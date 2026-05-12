package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var documentsListChildrenCmd = &cobra.Command{
	Use:   "list-children <job-id>",
	Short: "List child documents",
	Long:  `For container files (ZIP, TAR, EML), list the child documents extracted from the container.

Arguments:
  job-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel documents list-children <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		result, err := client.Documents.ListChildren(ctx, jobId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsListChildrenCmd)

}
