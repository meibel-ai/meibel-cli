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
	tagUpdateTagConfigData string
	tagUpdateTagConfigInteractive bool
)

var tagUpdateTagConfigCmd = &cobra.Command{
	Use:   "update-config <datasource-id>",
	Short: "Update Tag Config",
	Long:  `Update Tag Config

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources tag update-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.UpdateTagConfigRequest

		if tagUpdateTagConfigData != "" {
			if err := json.Unmarshal([]byte(tagUpdateTagConfigData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if tagUpdateTagConfigInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.Tag.UpdateTagConfig(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagUpdateTagConfigCmd)

	tagUpdateTagConfigCmd.Flags().StringVarP(&tagUpdateTagConfigData, "data", "d", "", "JSON data for the request body")
	tagUpdateTagConfigCmd.Flags().BoolVarP(&tagUpdateTagConfigInteractive, "interactive", "i", false, "use interactive form input")
}
