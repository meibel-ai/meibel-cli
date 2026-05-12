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
	datasourcesCreateData string
	datasourcesCreateInteractive bool
)

var datasourcesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Datasource",
	Long:  `Create Datasource`,
	Example: "meibel datasources create",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.CreateDatasourceRequest

		if datasourcesCreateData != "" {
			if err := json.Unmarshal([]byte(datasourcesCreateData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesCreateInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.Create(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesCmd.AddCommand(datasourcesCreateCmd)

	datasourcesCreateCmd.Flags().StringVarP(&datasourcesCreateData, "data", "d", "", "JSON data for the request body")
	datasourcesCreateCmd.Flags().BoolVarP(&datasourcesCreateInteractive, "interactive", "i", false, "use interactive form input")
}
