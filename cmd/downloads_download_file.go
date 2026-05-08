package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var downloadsDownloadFileCmd = &cobra.Command{
	Use:   "file <datasource-id> <job-id>",
	Short: "Download File",
	Long:  `Download File

Arguments:
  datasource-id: required
  job-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources downloads file <datasource-id> <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		jobId := args[1]

		result, err := client.Datasources.Downloads.DownloadFile(ctx, datasourceId, jobId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	downloadsCmd.AddCommand(downloadsDownloadFileCmd)

}
