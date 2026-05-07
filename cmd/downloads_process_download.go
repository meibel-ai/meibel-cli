package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	downloadsProcessDownloadData string
	downloadsProcessDownloadInteractive bool
)

var downloadsProcessDownloadCmd = &cobra.Command{
	Use:   "process <datasource-id>",
	Short: "Process Download (sync)",
	Long:  `Process Download (sync)

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources downloads process <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body interface{}

		if downloadsProcessDownloadData != "" {
			if err := json.Unmarshal([]byte(downloadsProcessDownloadData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Datasources.Downloads.ProcessDownload(ctx, datasourceId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	downloadsCmd.AddCommand(downloadsProcessDownloadCmd)

	downloadsProcessDownloadCmd.Flags().StringVarP(&downloadsProcessDownloadData, "data", "d", "", "JSON data for the request body")
	downloadsProcessDownloadCmd.Flags().BoolVarP(&downloadsProcessDownloadInteractive, "interactive", "i", false, "use interactive form input")
}
