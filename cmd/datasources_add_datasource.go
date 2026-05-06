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
	datasourcesAddDatasourceData string
	datasourcesAddDatasourceInteractive bool
)

var datasourcesAddDatasourceCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Datasource",
	Long:  `Add Datasource`,
	Example: "meibel datasources add",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.AddDatasourceRequest

		if datasourcesAddDatasourceData != "" {
			if err := json.Unmarshal([]byte(datasourcesAddDatasourceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesAddDatasourceInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("CustomerId").Description("").Value(&body.CustomerId),
					huh.NewInput().Title("ProjectId").Description("").Value(&body.ProjectId),
					huh.NewInput().Title("Name").Description("").Value(&body.Name),
					huh.NewInput().Title("Description").Description("").Value(&body.Description),
					huh.NewInput().Title("Recurrence").Description("").Value(&body.Recurrence),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Datasources.AddDatasource(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesAddDatasourceCmd)

	datasourcesAddDatasourceCmd.Flags().StringVarP(&datasourcesAddDatasourceData, "data", "d", "", "JSON data for the request body")
	datasourcesAddDatasourceCmd.Flags().BoolVarP(&datasourcesAddDatasourceInteractive, "interactive", "i", false, "use interactive form input")
}
