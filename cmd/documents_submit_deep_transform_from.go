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
	documentsSubmitDeepTransformFromData string
	documentsSubmitDeepTransformFromInteractive bool
)

var documentsSubmitDeepTransformFromCmd = &cobra.Command{
	Use:   "submit-deep-transform-from",
	Short: "Submit a deep-transform extraction reusing a parsed document (async)",
	Long:  `Submit an extraction that reuses an already-parsed document (by 'document_job_id' from POST /documents) instead of re-parsing an upload. Returns immediately with a job id. Poll status via GET /documents/deep-transform/{job_id} and download artifacts once it succeeds. Submission is idempotent on the (document, schema) pair.`,
	Example: "meibel documents submit-deep-transform-from",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.SubmitDeepTransformFromDocument

		if documentsSubmitDeepTransformFromData != "" {
			if err := json.Unmarshal([]byte(documentsSubmitDeepTransformFromData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if documentsSubmitDeepTransformFromInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("DocumentJobId").Description("A document job id returned by POST /documents. Reuses that parse so the document is not parsed again. The document must belong to the calling customer.").Value(&body.DocumentJobId),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.Documents.SubmitDeepTransformFrom(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsSubmitDeepTransformFromCmd)

	documentsSubmitDeepTransformFromCmd.Flags().StringVarP(&documentsSubmitDeepTransformFromData, "data", "d", "", "JSON data for the request body")
	documentsSubmitDeepTransformFromCmd.Flags().BoolVarP(&documentsSubmitDeepTransformFromInteractive, "interactive", "i", false, "use interactive form input")
}
