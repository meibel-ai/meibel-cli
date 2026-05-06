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
	ragAddRagConfigData string
	ragAddRagConfigInteractive bool
)

var ragAddRagConfigCmd = &cobra.Command{
	Use:   "add-config <datasource-id>",
	Short: "Add Rag Config",
	Long:  `Add Rag Config

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources rag add-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.AddRagConfigRequest

		if ragAddRagConfigData != "" {
			if err := json.Unmarshal([]byte(ragAddRagConfigData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if ragAddRagConfigInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("CollectionId").Description("").Value(&body.CollectionId),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Datasources.Rag.AddRagConfig(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	ragCmd.AddCommand(ragAddRagConfigCmd)

	ragAddRagConfigCmd.Flags().StringVarP(&ragAddRagConfigData, "data", "d", "", "JSON data for the request body")
	ragAddRagConfigCmd.Flags().BoolVarP(&ragAddRagConfigInteractive, "interactive", "i", false, "use interactive form input")
}
