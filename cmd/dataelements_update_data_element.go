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
	dataelementsUpdateDataElementData string
	dataelementsUpdateDataElementInteractive bool
)

var dataelementsUpdateDataElementCmd = &cobra.Command{
	Use:   "update-data-element <datasource-id> <data-element-id>",
	Short: "Update Data Element",
	Long:  `Update Data Element

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources dataelements update-data-element <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		var body sdk.UpdateDataElementRequest

		if dataelementsUpdateDataElementData != "" {
			if err := json.Unmarshal([]byte(dataelementsUpdateDataElementData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if dataelementsUpdateDataElementInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.Dataelements.UpdateDataElement(ctx, datasourceId, dataElementId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataelementsCmd.AddCommand(dataelementsUpdateDataElementCmd)

	dataelementsUpdateDataElementCmd.Flags().StringVarP(&dataelementsUpdateDataElementData, "data", "d", "", "JSON data for the request body")
	dataelementsUpdateDataElementCmd.Flags().BoolVarP(&dataelementsUpdateDataElementInteractive, "interactive", "i", false, "use interactive form input")
}
