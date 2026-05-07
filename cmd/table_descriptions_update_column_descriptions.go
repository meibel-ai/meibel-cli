package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	tableDescriptionsUpdateColumnDescriptionsData string
	tableDescriptionsUpdateColumnDescriptionsInteractive bool
)

var tableDescriptionsUpdateColumnDescriptionsCmd = &cobra.Command{
	Use:   "update-column <datasource-id> <table-name>",
	Short: "Update Column Descriptions",
	Long:  `Update Column Descriptions

Arguments:
  datasource-id: required
  table-name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources table-descriptions update-column <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		var body []TagColumnUpdateItem

		if tableDescriptionsUpdateColumnDescriptionsData != "" {
			if err := json.Unmarshal([]byte(tableDescriptionsUpdateColumnDescriptionsData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Datasources.TableDescriptions.UpdateColumnDescriptions(ctx, datasourceId, tableName, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tableDescriptionsCmd.AddCommand(tableDescriptionsUpdateColumnDescriptionsCmd)

	tableDescriptionsUpdateColumnDescriptionsCmd.Flags().StringVarP(&tableDescriptionsUpdateColumnDescriptionsData, "data", "d", "", "JSON data for the request body")
	tableDescriptionsUpdateColumnDescriptionsCmd.Flags().BoolVarP(&tableDescriptionsUpdateColumnDescriptionsInteractive, "interactive", "i", false, "use interactive form input")
}
