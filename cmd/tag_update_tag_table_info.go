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
	tagUpdateTagTableInfoData string
	tagUpdateTagTableInfoInteractive bool
)

var tagUpdateTagTableInfoCmd = &cobra.Command{
	Use:   "update-table-info <datasource-id> <table-name>",
	Short: "Update Tag Table Info",
	Long:  `Update Tag Table Info

Arguments:
  datasource-id: required
  table-name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources tag update-table-info <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		var body sdk.UpdateTagTableRequest

		if tagUpdateTagTableInfoData != "" {
			if err := json.Unmarshal([]byte(tagUpdateTagTableInfoData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if tagUpdateTagTableInfoInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.Tag.UpdateTagTableInfo(ctx, datasourceId, tableName, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagUpdateTagTableInfoCmd)

	tagUpdateTagTableInfoCmd.Flags().StringVarP(&tagUpdateTagTableInfoData, "data", "d", "", "JSON data for the request body")
	tagUpdateTagTableInfoCmd.Flags().BoolVarP(&tagUpdateTagTableInfoInteractive, "interactive", "i", false, "use interactive form input")
}
