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
	datasourcesUpdateData string
	datasourcesUpdateInteractive bool
)

var datasourcesUpdateCmd = &cobra.Command{
	Use:   "update <datasource-id>",
	Short: "Update Datasource",
	Long:  `Update Datasource

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources update <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.UpdateDatasourceRequest

		if datasourcesUpdateData != "" {
			if err := json.Unmarshal([]byte(datasourcesUpdateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesUpdateInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.Update(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesUpdateCmd)

	datasourcesUpdateCmd.Flags().StringVarP(&datasourcesUpdateData, "data", "d", "", "JSON data for the request body")
	datasourcesUpdateCmd.Flags().BoolVarP(&datasourcesUpdateInteractive, "interactive", "i", false, "use interactive form input")
}
