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
	dataElementsUpdateData string
	dataElementsUpdateInteractive bool
)

var dataElementsUpdateCmd = &cobra.Command{
	Use:   "update <data-element-id> <datasource-id>",
	Short: "Update Data Element",
	Long:  `Update Data Element

Arguments:
  data-element-id: required
  datasource-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources data-elements update <data-element-id> <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		dataElementId := args[0]
		datasourceId := args[1]

		var body sdk.UpdateDataElementRequest

		if dataElementsUpdateData != "" {
			if err := json.Unmarshal([]byte(dataElementsUpdateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if dataElementsUpdateInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.DataElements.Update(ctx, dataElementId, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementsCmd.AddCommand(dataElementsUpdateCmd)

	dataElementsUpdateCmd.Flags().StringVarP(&dataElementsUpdateData, "data", "d", "", "JSON data for the request body")
	dataElementsUpdateCmd.Flags().BoolVarP(&dataElementsUpdateInteractive, "interactive", "i", false, "use interactive form input")
}
