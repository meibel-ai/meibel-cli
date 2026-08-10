package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/config"
	"github.com/meibel-ai/meibel-cli/internal/tui"
	"github.com/meibel-ai/meibel-cli/internal/pathutil"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	documentsTransformFile string
	documentsTransformSchema string
	documentsTransformModel string
	documentsTransformPrompt string
	documentsTransformPromptId string
	documentsTransformTimeoutSeconds int64
	documentsTransformTrace bool
	documentsTransformBrowser bool
	documentsTransformWait bool
)

var documentsTransformCmd = &cobra.Command{
	Use:   "transform",
	Short: "Transform a document using AI extraction (sync)",
	Long:  `Upload a document for AI-powered structured extraction and block until complete. The file is uploaded to cloud storage and processed by a system agent.`,
	Example: "meibel documents transform",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if documentsTransformFile == "" && term.IsTerminal(int(os.Stdin.Fd())) {
			if err := huh.NewForm(huh.NewGroup(
				huh.NewInput().
					Title("File path").
					Description("Paste or type a path — leave blank to browse").
					Value(&documentsTransformFile),
			)).Run(); err != nil {
				return err
			}
			documentsTransformFile = pathutil.Expand(documentsTransformFile)
		}

		if documentsTransformFile == "" {
			picker := huh.NewFilePicker().
				Title("Select a file").
				CurrentDirectory(pathutil.StartDir()).
				FileAllowed(true).
				DirAllowed(false).
				ShowHidden(false).
				ShowSize(true).
				ShowPermissions(false).
				Height(15).
				Value(&documentsTransformFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if documentsTransformFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		opts := sdk.DocumentsTransformOptions{}
		opts.File = documentsTransformFile
		opts.Schema = documentsTransformSchema
		if documentsTransformModel != "" {
			opts.Model = &documentsTransformModel
		}
		if documentsTransformPrompt != "" {
			opts.Prompt = &documentsTransformPrompt
		}
		if documentsTransformPromptId != "" {
			opts.PromptId = &documentsTransformPromptId
		}
		if documentsTransformTimeoutSeconds != 0 {
			opts.TimeoutSeconds = &documentsTransformTimeoutSeconds
		}

		if documentsTransformWait {
			sp := tui.StartSpinner("Transform a document using AI extraction (sync)")
			result, err := client.Documents.SubmitDeepTransform(ctx, sdk.DocumentsSubmitDeepTransformOptions{File: documentsTransformFile, Schema: documentsTransformSchema})
			sp.Stop()
			if err != nil {
				return err
			}
			return output.Print(result)
		}

		sp := tui.StartSpinner("Transform a document using AI extraction (sync)")
		result, err := client.Documents.Transform(ctx, opts)
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

		if documentsTransformBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if documentsTransformTrace && jr.JobID != "" {
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
	documentsCmd.AddCommand(documentsTransformCmd)

	documentsTransformCmd.Flags().StringVarP(&documentsTransformFile, "file", "f", "", "Document file to transform")
	documentsTransformCmd.MarkFlagFilename("file")
	documentsTransformCmd.Flags().StringVar(&documentsTransformSchema, "schema", "", "JSON Schema dict (as JSON string) or schema name/ID")
	documentsTransformCmd.MarkFlagRequired("schema")
	documentsTransformCmd.Flags().StringVar(&documentsTransformModel, "model", "", "LLM model override")
	documentsTransformCmd.Flags().StringVar(&documentsTransformPrompt, "prompt", "", "Extraction instructions override")
	documentsTransformCmd.Flags().StringVar(&documentsTransformPromptId, "prompt-id", "", "Prompt template reference")
	documentsTransformCmd.Flags().Int64Var(&documentsTransformTimeoutSeconds, "timeout-seconds", 0, "Max wait time in seconds (sync only)")
	documentsTransformCmd.Flags().BoolVar(&documentsTransformTrace, "trace", false, "stream parsing trace after upload")
	documentsTransformCmd.Flags().BoolVar(&documentsTransformBrowser, "browser", false, "open trace in console")
	documentsTransformCmd.Flags().BoolVar(&documentsTransformWait, "wait", false, "wait for completion (synchronous)")
}
