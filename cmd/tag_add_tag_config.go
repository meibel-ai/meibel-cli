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
	tagAddTagConfigData string
	tagAddTagConfigInteractive bool
)

var tagAddTagConfigCmd = &cobra.Command{
	Use:   "add-config <datasource-id>",
	Short: "Add Tag Config",
	Long:  `Add Tag Config

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources tag add-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.AddTagConfigRequest

		if tagAddTagConfigData != "" {
			if err := json.Unmarshal([]byte(tagAddTagConfigData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if tagAddTagConfigInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("WorkingBucket").Description("").Value(&body.WorkingBucket),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Datasources.Tag.AddTagConfig(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagCmd.AddCommand(tagAddTagConfigCmd)

	tagAddTagConfigCmd.Flags().StringVarP(&tagAddTagConfigData, "data", "d", "", "JSON data for the request body")
	tagAddTagConfigCmd.Flags().BoolVarP(&tagAddTagConfigInteractive, "interactive", "i", false, "use interactive form input")
}
