package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var documentsGetDeepTransformStatusCmd = &cobra.Command{
	Use:   "get-deep-transform-status <job-id>",
	Short: "Get deep-transform job status",
	Long:  `Check status and, once succeeded, the list of downloadable artifacts.

Arguments:
  job-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel documents get-deep-transform-status <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		result, err := client.Documents.GetDeepTransformStatus(ctx, jobId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsGetDeepTransformStatusCmd)

}
