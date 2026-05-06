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
	ragUpdateChunkingStrategyData string
	ragUpdateChunkingStrategyInteractive bool
)

var ragUpdateChunkingStrategyCmd = &cobra.Command{
	Use:   "update-chunking-strategy <datasource-id>",
	Short: "Update Chunking Strategy",
	Long:  `Update Chunking Strategy

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources rag update-chunking-strategy <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.UpdateChunkingStrategyRequest

		if ragUpdateChunkingStrategyData != "" {
			if err := json.Unmarshal([]byte(ragUpdateChunkingStrategyData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if ragUpdateChunkingStrategyInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Datasources.Rag.UpdateChunkingStrategy(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	ragCmd.AddCommand(ragUpdateChunkingStrategyCmd)

	ragUpdateChunkingStrategyCmd.Flags().StringVarP(&ragUpdateChunkingStrategyData, "data", "d", "", "JSON data for the request body")
	ragUpdateChunkingStrategyCmd.Flags().BoolVarP(&ragUpdateChunkingStrategyInteractive, "interactive", "i", false, "use interactive form input")
}
