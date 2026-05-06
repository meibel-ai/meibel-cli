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
	dataelementsGetDataElementsByFiltersRegexFilter string
	dataelementsGetDataElementsByFiltersMediaTypeFilters []string
	dataelementsGetDataElementsByFiltersOffset int64
	dataelementsGetDataElementsByFiltersLimit int64
	dataelementsGetDataElementsByFiltersSortBy string
	dataelementsGetDataElementsByFiltersSortOrder string
	dataelementsGetDataElementsByFiltersData string
	dataelementsGetDataElementsByFiltersInteractive bool
)

var dataelementsGetDataElementsByFiltersCmd = &cobra.Command{
	Use:   "get-data-elements-by-filters <datasource-id>",
	Short: "Get Data Elements By Filters",
	Long:  `Get Data Elements By Filters

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources dataelements get-data-elements-by-filters <datasource-id> --regex-filter=<value> --media-type-filters=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.DataElementFilterRequest

		if dataelementsGetDataElementsByFiltersData != "" {
			if err := json.Unmarshal([]byte(dataelementsGetDataElementsByFiltersData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if dataelementsGetDataElementsByFiltersInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		opts := &sdk.GetDataElementsByFiltersOptions{}
		if dataelementsGetDataElementsByFiltersRegexFilter != "" {
			opts.RegexFilter = &dataelementsGetDataElementsByFiltersRegexFilter
		}
		if dataelementsGetDataElementsByFiltersOffset != 0 {
			opts.Offset = &dataelementsGetDataElementsByFiltersOffset
		}
		if dataelementsGetDataElementsByFiltersLimit != 0 {
			opts.Limit = &dataelementsGetDataElementsByFiltersLimit
		}
		if dataelementsGetDataElementsByFiltersSortBy != "" {
			opts.SortBy = &dataelementsGetDataElementsByFiltersSortBy
		}
		if dataelementsGetDataElementsByFiltersSortOrder != "" {
			opts.SortOrder = &dataelementsGetDataElementsByFiltersSortOrder
		}

		result, err := client.Datasources.Dataelements.GetDataElementsByFilters(ctx, datasourceId, &body, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataelementsCmd.AddCommand(dataelementsGetDataElementsByFiltersCmd)

	dataelementsGetDataElementsByFiltersCmd.Flags().StringVarP(&dataelementsGetDataElementsByFiltersRegexFilter, "regex-filter", "", "", "The regex-filter parameter")
	dataelementsGetDataElementsByFiltersCmd.Flags().StringSliceVarP(&dataelementsGetDataElementsByFiltersMediaTypeFilters, "media-type-filters", "", nil, "The media-type-filters parameter")
	dataelementsGetDataElementsByFiltersCmd.Flags().Int64VarP(&dataelementsGetDataElementsByFiltersOffset, "offset", "", 0, "Number of items to skip")
	dataelementsGetDataElementsByFiltersCmd.Flags().Int64VarP(&dataelementsGetDataElementsByFiltersLimit, "limit", "", 10, "Maximum number of items to return")
	dataelementsGetDataElementsByFiltersCmd.Flags().StringVarP(&dataelementsGetDataElementsByFiltersSortBy, "sort-by", "", "", "Field to sort by")
	dataelementsGetDataElementsByFiltersCmd.Flags().StringVarP(&dataelementsGetDataElementsByFiltersSortOrder, "sort-order", "", "", "Sort order (asc or desc)")
	dataelementsGetDataElementsByFiltersCmd.Flags().StringVarP(&dataelementsGetDataElementsByFiltersData, "data", "d", "", "JSON data for the request body")
	dataelementsGetDataElementsByFiltersCmd.Flags().BoolVarP(&dataelementsGetDataElementsByFiltersInteractive, "interactive", "i", false, "use interactive form input")
}
