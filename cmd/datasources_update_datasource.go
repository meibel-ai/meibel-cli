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
	datasourcesUpdateDatasourceData string
	datasourcesUpdateDatasourceInteractive bool
)

var datasourcesUpdateDatasourceCmd = &cobra.Command{
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

		if datasourcesUpdateDatasourceData != "" {
			if err := json.Unmarshal([]byte(datasourcesUpdateDatasourceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesUpdateDatasourceInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.UpdateDatasource(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesUpdateDatasourceCmd)

	datasourcesUpdateDatasourceCmd.Flags().StringVarP(&datasourcesUpdateDatasourceData, "data", "d", "", "JSON data for the request body")
	datasourcesUpdateDatasourceCmd.Flags().BoolVarP(&datasourcesUpdateDatasourceInteractive, "interactive", "i", false, "use interactive form input")
}
