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
	dataElementsUpdateDataElementData string
	dataElementsUpdateDataElementInteractive bool
)

var dataElementsUpdateDataElementCmd = &cobra.Command{
	Use:   "update <datasource-id> <data-element-id>",
	Short: "Update Data Element",
	Long:  `Update Data Element

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources data-elements update <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		var body sdk.UpdateDataElementRequest

		if dataElementsUpdateDataElementData != "" {
			if err := json.Unmarshal([]byte(dataElementsUpdateDataElementData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if dataElementsUpdateDataElementInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.DataElements.UpdateDataElement(ctx, datasourceId, dataElementId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementsCmd.AddCommand(dataElementsUpdateDataElementCmd)

	dataElementsUpdateDataElementCmd.Flags().StringVarP(&dataElementsUpdateDataElementData, "data", "d", "", "JSON data for the request body")
	dataElementsUpdateDataElementCmd.Flags().BoolVarP(&dataElementsUpdateDataElementInteractive, "interactive", "i", false, "use interactive form input")
}
