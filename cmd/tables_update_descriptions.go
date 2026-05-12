package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
)

var (
	tablesUpdateDescriptionsData string
	tablesUpdateDescriptionsInteractive bool
)

var tablesUpdateDescriptionsCmd = &cobra.Command{
	Use:   "update-descriptions <datasource-id>",
	Short: "Update Table Descriptions",
	Long:  `Update Table Descriptions

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources tables update-descriptions <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body []TagTableUpdateItem

		if tablesUpdateDescriptionsData != "" {
			if err := json.Unmarshal([]byte(tablesUpdateDescriptionsData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.Datasources.Tables.UpdateDescriptions(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tablesCmd.AddCommand(tablesUpdateDescriptionsCmd)

	tablesUpdateDescriptionsCmd.Flags().StringVarP(&tablesUpdateDescriptionsData, "data", "d", "", "JSON data for the request body")
	tablesUpdateDescriptionsCmd.Flags().BoolVarP(&tablesUpdateDescriptionsInteractive, "interactive", "i", false, "use interactive form input")
}
