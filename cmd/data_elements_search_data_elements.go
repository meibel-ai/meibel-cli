package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	dataElementsSearchDataElementsCursor string
	dataElementsSearchDataElementsLimit int64
	dataElementsSearchDataElementsData string
	dataElementsSearchDataElementsInteractive bool
)

var dataElementsSearchDataElementsCmd = &cobra.Command{
	Use:   "search <datasource-id>",
	Short: "Search Data Elements",
	Long:  `Search Data Elements

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel data-elements search <datasource-id> --cursor=<value> --limit=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.DataElementSearchRequest

		if dataElementsSearchDataElementsData != "" {
			if err := json.Unmarshal([]byte(dataElementsSearchDataElementsData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if dataElementsSearchDataElementsInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		opts := &sdk.SearchDataElementsOptions{}
		if dataElementsSearchDataElementsCursor != "" {
			opts.Cursor = &dataElementsSearchDataElementsCursor
		}
		if dataElementsSearchDataElementsLimit != 0 {
			opts.Limit = &dataElementsSearchDataElementsLimit
		}

		result, err := client.DataElements.SearchDataElements(ctx, datasourceId, body, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementsCmd.AddCommand(dataElementsSearchDataElementsCmd)

	dataElementsSearchDataElementsCmd.Flags().StringVarP(&dataElementsSearchDataElementsCursor, "cursor", "", "", "Cursor for pagination")
	dataElementsSearchDataElementsCmd.Flags().Int64VarP(&dataElementsSearchDataElementsLimit, "limit", "", 100, "Maximum items to return")
	dataElementsSearchDataElementsCmd.Flags().StringVarP(&dataElementsSearchDataElementsData, "data", "d", "", "JSON data for the request body")
	dataElementsSearchDataElementsCmd.Flags().BoolVarP(&dataElementsSearchDataElementsInteractive, "interactive", "i", false, "use interactive form input")
}
