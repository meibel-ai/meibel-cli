package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	tablesUpdateColumnDescriptionsData string
	tablesUpdateColumnDescriptionsInteractive bool
)

var tablesUpdateColumnDescriptionsCmd = &cobra.Command{
	Use:   "update-column-descriptions <table-name> <datasource-id>",
	Short: "Update Column Descriptions",
	Long:  `Update Column Descriptions

Arguments:
  table-name: required
  datasource-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources tables update-column-descriptions <table-name> <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		tableName := args[0]
		datasourceId := args[1]

		var body sdk.UpdateTagColumnsRequest

		if tablesUpdateColumnDescriptionsData != "" {
			if err := json.Unmarshal([]byte(tablesUpdateColumnDescriptionsData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if tablesUpdateColumnDescriptionsInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Datasources.Tables.UpdateColumnDescriptions(ctx, tableName, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tablesCmd.AddCommand(tablesUpdateColumnDescriptionsCmd)

	tablesUpdateColumnDescriptionsCmd.Flags().StringVarP(&tablesUpdateColumnDescriptionsData, "data", "d", "", "JSON data for the request body")
	tablesUpdateColumnDescriptionsCmd.Flags().BoolVarP(&tablesUpdateColumnDescriptionsInteractive, "interactive", "i", false, "use interactive form input")
}
