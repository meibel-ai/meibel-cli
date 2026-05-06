package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	tagUpdateTagColumnInfoData string
	tagUpdateTagColumnInfoInteractive bool
)

var tagUpdateTagColumnInfoCmd = &cobra.Command{
	Use:   "update-column-info <datasource-id> <table-name> <column-name>",
	Short: "Update Tag Column Info",
	Long:  `Update Tag Column Info

Arguments:
  datasource-id: required
  table-name: required
  column-name: required`,
	Args:  cobra.ExactArgs(3),
	Example: "meibel datasources tag update-column-info <datasource-id> <table-name> <column-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]
		columnName := args[2]

		var body sdk.UpdateTagColumnRequest

		if tagUpdateTagColumnInfoData != "" {
			if err := json.Unmarshal([]byte(tagUpdateTagColumnInfoData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if tagUpdateTagColumnInfoInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.Tag.UpdateTagColumnInfo(ctx, datasourceId, tableName, columnName, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagUpdateTagColumnInfoCmd)

	tagUpdateTagColumnInfoCmd.Flags().StringVarP(&tagUpdateTagColumnInfoData, "data", "d", "", "JSON data for the request body")
	tagUpdateTagColumnInfoCmd.Flags().BoolVarP(&tagUpdateTagColumnInfoInteractive, "interactive", "i", false, "use interactive form input")
}
