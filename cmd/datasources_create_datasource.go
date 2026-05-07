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
	datasourcesCreateDatasourceData string
	datasourcesCreateDatasourceInteractive bool
)

var datasourcesCreateDatasourceCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Datasource",
	Long:  `Create Datasource`,
	Example: "meibel datasources create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateDatasourceRequest

		if datasourcesCreateDatasourceData != "" {
			if err := json.Unmarshal([]byte(datasourcesCreateDatasourceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesCreateDatasourceInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Name").Description("Human-readable datasource name").Value(&body.Name),
					huh.NewInput().Title("Description").Description("What this datasource contains"),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Datasources.CreateDatasource(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesCreateDatasourceCmd)

	datasourcesCreateDatasourceCmd.Flags().StringVarP(&datasourcesCreateDatasourceData, "data", "d", "", "JSON data for the request body")
	datasourcesCreateDatasourceCmd.Flags().BoolVarP(&datasourcesCreateDatasourceInteractive, "interactive", "i", false, "use interactive form input")
}
