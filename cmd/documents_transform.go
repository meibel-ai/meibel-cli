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
	documentsTransformData string
	documentsTransformInteractive bool
)

var documentsTransformCmd = &cobra.Command{
	Use:   "transform",
	Short: "Transform a document using AI extraction (sync)",
	Long:  `Submit a document for AI-powered structured extraction and block until complete. Internally orchestrates a system agent session, polls for completion, and returns the extracted data.`,
	Example: "meibel documents transform",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.TransformDocumentRequest

		if documentsTransformData != "" {
			if err := json.Unmarshal([]byte(documentsTransformData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if documentsTransformInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("File").Description("File path, URL, or GCS URI to transform").Value(&body.File),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Documents.Transform(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsTransformCmd)

	documentsTransformCmd.Flags().StringVarP(&documentsTransformData, "data", "d", "", "JSON data for the request body")
	documentsTransformCmd.Flags().BoolVarP(&documentsTransformInteractive, "interactive", "i", false, "use interactive form input")
}
