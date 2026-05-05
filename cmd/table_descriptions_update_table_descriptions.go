package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel/internal/output"
)

var (
	tableDescriptionsUpdateTableDescriptionsData string
	tableDescriptionsUpdateTableDescriptionsInteractive bool
)

var tableDescriptionsUpdateTableDescriptionsCmd = &cobra.Command{
	Use:   "update <datasource-id>",
	Short: "Update Table Descriptions",
	Long:  `Update Table Descriptions

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel table-descriptions update <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body []TagTableUpdateItem

		if tableDescriptionsUpdateTableDescriptionsData != "" {
			if err := json.Unmarshal([]byte(tableDescriptionsUpdateTableDescriptionsData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.TableDescriptions.UpdateTableDescriptions(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tableDescriptionsCmd.AddCommand(tableDescriptionsUpdateTableDescriptionsCmd)

	tableDescriptionsUpdateTableDescriptionsCmd.Flags().StringVarP(&tableDescriptionsUpdateTableDescriptionsData, "data", "d", "", "JSON data for the request body")
	tableDescriptionsUpdateTableDescriptionsCmd.Flags().BoolVarP(&tableDescriptionsUpdateTableDescriptionsInteractive, "interactive", "i", false, "use interactive form input")
}
