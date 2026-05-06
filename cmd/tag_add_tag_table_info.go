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
	tagAddTagTableInfoData string
	tagAddTagTableInfoInteractive bool
)

var tagAddTagTableInfoCmd = &cobra.Command{
	Use:   "add-table-info <datasource-id> <table-name>",
	Short: "Add Tag Table Info",
	Long:  `Add Tag Table Info

Arguments:
  datasource-id: required
  table-name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources tag add-table-info <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		var body sdk.AddTagTableRequest

		if tagAddTagTableInfoData != "" {
			if err := json.Unmarshal([]byte(tagAddTagTableInfoData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if tagAddTagTableInfoInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.Tag.AddTagTableInfo(ctx, datasourceId, tableName, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagAddTagTableInfoCmd)

	tagAddTagTableInfoCmd.Flags().StringVarP(&tagAddTagTableInfoData, "data", "d", "", "JSON data for the request body")
	tagAddTagTableInfoCmd.Flags().BoolVarP(&tagAddTagTableInfoInteractive, "interactive", "i", false, "use interactive form input")
}
