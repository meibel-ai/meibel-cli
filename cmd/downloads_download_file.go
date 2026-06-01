package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var downloadsDownloadFileCmd = &cobra.Command{
	Use:   "file <job-id> <datasource-id>",
	Short: "Download File",
	Long:  `Download File

Arguments:
  job-id: required
  datasource-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources downloads file <job-id> <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]
		datasourceId := args[1]

		result, err := client.Datasources.Downloads.DownloadFile(ctx, jobId, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	downloadsCmd.AddCommand(downloadsDownloadFileCmd)

}
