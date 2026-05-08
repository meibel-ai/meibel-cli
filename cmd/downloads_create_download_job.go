package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var (
	downloadsCreateDownloadJobData string
	downloadsCreateDownloadJobInteractive bool
)

var downloadsCreateDownloadJobCmd = &cobra.Command{
	Use:   "create-job <datasource-id>",
	Short: "Create Download Job (async)",
	Long:  `Create Download Job (async)

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources downloads create-job <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body interface{}

		if downloadsCreateDownloadJobData != "" {
			if err := json.Unmarshal([]byte(downloadsCreateDownloadJobData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Datasources.Downloads.CreateDownloadJob(ctx, datasourceId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	downloadsCmd.AddCommand(downloadsCreateDownloadJobCmd)

	downloadsCreateDownloadJobCmd.Flags().StringVarP(&downloadsCreateDownloadJobData, "data", "d", "", "JSON data for the request body")
	downloadsCreateDownloadJobCmd.Flags().BoolVarP(&downloadsCreateDownloadJobInteractive, "interactive", "i", false, "use interactive form input")
}
