package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/config"
	"github.com/meibel-ai/meibel-cli/internal/tui"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	documentsSubmitTransformFile string
	documentsSubmitTransformSchema string
	documentsSubmitTransformModel string
	documentsSubmitTransformPrompt string
	documentsSubmitTransformPromptId string
	documentsSubmitTransformTimeoutSeconds int64
	documentsSubmitTransformTrace bool
	documentsSubmitTransformBrowser bool
)

var documentsSubmitTransformCmd = &cobra.Command{
	Use:   "submit-transform",
	Short: "Submit a document transform (async)",
	Long:  `Upload a document for AI-powered extraction and return immediately. Poll for completion via client.sessions.get(execution_id).`,
	Example: "meibel documents submit-transform",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if documentsSubmitTransformFile == "" {
			home, _ := os.UserHomeDir()
			if home == "" {
				home, _ = os.Getwd()
			}
			picker := huh.NewFilePicker().
				Title("Select a file").
				CurrentDirectory(home).
				FileAllowed(true).
				DirAllowed(false).
				ShowHidden(false).
				ShowSize(true).
				ShowPermissions(false).
				Height(15).
				Value(&documentsSubmitTransformFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if documentsSubmitTransformFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		opts := sdk.DocumentsSubmitTransformOptions{}
		opts.File = documentsSubmitTransformFile
		opts.Schema = documentsSubmitTransformSchema
		if documentsSubmitTransformModel != "" {
			opts.Model = &documentsSubmitTransformModel
		}
		if documentsSubmitTransformPrompt != "" {
			opts.Prompt = &documentsSubmitTransformPrompt
		}
		if documentsSubmitTransformPromptId != "" {
			opts.PromptId = &documentsSubmitTransformPromptId
		}
		if documentsSubmitTransformTimeoutSeconds != 0 {
			opts.TimeoutSeconds = &documentsSubmitTransformTimeoutSeconds
		}

		sp := tui.StartSpinner("Submit a document transform (async)")
		result, err := client.Documents.SubmitTransform(ctx, opts)
		sp.Stop()
		if err != nil {
			return err
		}

		type jobResult struct {
			JobID string `json:"job_id"`
		}
		var jr jobResult
		b, _ := json.Marshal(result)
		json.Unmarshal(b, &jr)

		if documentsSubmitTransformBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if documentsSubmitTransformTrace && jr.JobID != "" {
			output.Print(result)

			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
			defer cancel()

			stream, err := client.Documents.StreamTrace(ctx, jr.JobID)
			if err != nil {
				return err
			}
			defer stream.Close()

			return tui.StreamEvents(ctx, stream)
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsSubmitTransformCmd)

	documentsSubmitTransformCmd.Flags().StringVarP(&documentsSubmitTransformFile, "file", "f", "", "Document file to transform")
	documentsSubmitTransformCmd.MarkFlagFilename("file")
	documentsSubmitTransformCmd.Flags().StringVar(&documentsSubmitTransformSchema, "schema", "", "JSON Schema dict (as JSON string) or schema name/ID")
	documentsSubmitTransformCmd.MarkFlagRequired("schema")
	documentsSubmitTransformCmd.Flags().StringVar(&documentsSubmitTransformModel, "model", "", "LLM model override")
	documentsSubmitTransformCmd.Flags().StringVar(&documentsSubmitTransformPrompt, "prompt", "", "Extraction instructions override")
	documentsSubmitTransformCmd.Flags().StringVar(&documentsSubmitTransformPromptId, "prompt-id", "", "Prompt template reference")
	documentsSubmitTransformCmd.Flags().Int64Var(&documentsSubmitTransformTimeoutSeconds, "timeout-seconds", 0, "Max wait time in seconds (sync only)")
	documentsSubmitTransformCmd.Flags().BoolVar(&documentsSubmitTransformTrace, "trace", false, "stream parsing trace after upload")
	documentsSubmitTransformCmd.Flags().BoolVar(&documentsSubmitTransformBrowser, "browser", false, "open trace in console")
}
