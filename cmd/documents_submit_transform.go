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
	documentsSubmitTransformData string
	documentsSubmitTransformInteractive bool
)

var documentsSubmitTransformCmd = &cobra.Command{
	Use:   "submit-transform",
	Short: "Submit a document transform (async)",
	Long:  `Submit a document for AI-powered extraction and return immediately. Poll for completion via client.sessions.get(execution_id).`,
	Example: "meibel documents submit-transform",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.TransformDocumentRequest

		if documentsSubmitTransformData != "" {
			if err := json.Unmarshal([]byte(documentsSubmitTransformData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if documentsSubmitTransformInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.Documents.SubmitTransform(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsSubmitTransformCmd)

	documentsSubmitTransformCmd.Flags().StringVarP(&documentsSubmitTransformData, "data", "d", "", "JSON data for the request body")
	documentsSubmitTransformCmd.Flags().BoolVarP(&documentsSubmitTransformInteractive, "interactive", "i", false, "use interactive form input")
}
