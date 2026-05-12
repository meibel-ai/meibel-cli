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
	dataElementsSearchCursor string
	dataElementsSearchLimit int64
	dataElementsSearchData string
	dataElementsSearchInteractive bool
)

var dataElementsSearchCmd = &cobra.Command{
	Use:   "search <datasource-id>",
	Short: "Search Data Elements",
	Long:  `Search Data Elements

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources data-elements search <datasource-id> --cursor=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.DataElementSearchRequest

		if dataElementsSearchData != "" {
			if err := json.Unmarshal([]byte(dataElementsSearchData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if dataElementsSearchInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		opts := &sdk.DataElementsSearchOptions{}
		if dataElementsSearchCursor != "" {
			opts.Cursor = &dataElementsSearchCursor
		}
		if dataElementsSearchLimit != 0 {
			opts.Limit = &dataElementsSearchLimit
		}

		result, err := client.Datasources.DataElements.Search(ctx, datasourceId, body, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementsCmd.AddCommand(dataElementsSearchCmd)

	dataElementsSearchCmd.Flags().StringVarP(&dataElementsSearchCursor, "cursor", "", "", "Cursor for pagination")
	dataElementsSearchCmd.Flags().Int64VarP(&dataElementsSearchLimit, "limit", "", 100, "Maximum items to return")
	dataElementsSearchCmd.Flags().StringVarP(&dataElementsSearchData, "data", "d", "", "JSON data for the request body")
	dataElementsSearchCmd.Flags().BoolVarP(&dataElementsSearchInteractive, "interactive", "i", false, "use interactive form input")
}
