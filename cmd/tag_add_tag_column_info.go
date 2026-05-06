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
	tagAddTagColumnInfoData string
	tagAddTagColumnInfoInteractive bool
)

var tagAddTagColumnInfoCmd = &cobra.Command{
	Use:   "add-column-info <datasource-id> <table-name> <column-name>",
	Short: "Add Tag Column Info",
	Long:  `Add Tag Column Info

Arguments:
  datasource-id: required
  table-name: required
  column-name: required`,
	Args:  cobra.ExactArgs(3),
	Example: "meibel datasources tag add-column-info <datasource-id> <table-name> <column-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]
		columnName := args[2]

		var body sdk.AddTagColumnRequest

		if tagAddTagColumnInfoData != "" {
			if err := json.Unmarshal([]byte(tagAddTagColumnInfoData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if tagAddTagColumnInfoInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.Tag.AddTagColumnInfo(ctx, datasourceId, tableName, columnName, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagAddTagColumnInfoCmd)

	tagAddTagColumnInfoCmd.Flags().StringVarP(&tagAddTagColumnInfoData, "data", "d", "", "JSON data for the request body")
	tagAddTagColumnInfoCmd.Flags().BoolVarP(&tagAddTagColumnInfoInteractive, "interactive", "i", false, "use interactive form input")
}
