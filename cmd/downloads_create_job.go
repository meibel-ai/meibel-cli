package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var (
	downloadsCreateJobData string
	downloadsCreateJobInteractive bool
)

var downloadsCreateJobCmd = &cobra.Command{
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

		if downloadsCreateJobData != "" {
			if err := json.Unmarshal([]byte(downloadsCreateJobData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Datasources.Downloads.CreateJob(ctx, datasourceId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	downloadsCmd.AddCommand(downloadsCreateJobCmd)

	downloadsCreateJobCmd.Flags().StringVarP(&downloadsCreateJobData, "data", "d", "", "JSON data for the request body")
	downloadsCreateJobCmd.Flags().BoolVarP(&downloadsCreateJobInteractive, "interactive", "i", false, "use interactive form input")
}
