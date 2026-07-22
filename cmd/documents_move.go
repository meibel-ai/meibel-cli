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
	documentsMoveData string
	documentsMoveInteractive bool
)

var documentsMoveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move documents into a datasource (async)",
	Long:  `Move documents (identified by their parse job IDs, e.g. the job_id returned by parseDocument) into an existing datasource or a newly created one. Returns a workflow_id to poll for completion.`,
	Example: "meibel documents move",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.MoveDocumentsRequest

		if documentsMoveData != "" {
			if err := json.Unmarshal([]byte(documentsMoveData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if documentsMoveInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Documents.Move(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsMoveCmd)

	documentsMoveCmd.Flags().StringVarP(&documentsMoveData, "data", "d", "", "JSON data for the request body")
	documentsMoveCmd.Flags().BoolVarP(&documentsMoveInteractive, "interactive", "i", false, "use interactive form input")
}
