package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"github.com/meibel-ai/meibel-go/meibel/internal/output"
	"github.com/meibel-ai/meibel-go/meibel/internal/config"
	"github.com/meibel-ai/meibel-go/meibel/internal/tui"
	"github.com/meibel-ai/meibel-go/meibel/internal/upload"
)

var (
	documentsSubmitTransformFile string
	documentsSubmitTransformArtifactSchema string
	documentsSubmitTransformModel string
	documentsSubmitTransformPrompt string
	documentsSubmitTransformPromptId string
	documentsSubmitTransformTimeoutSeconds string
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

		f, err := os.Open(documentsSubmitTransformFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(documentsSubmitTransformFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		result, err := client.Documents.SubmitTransform(ctx, pr, fileName, documentsSubmitTransformArtifactSchema, documentsSubmitTransformModel, documentsSubmitTransformPrompt, documentsSubmitTransformPromptId, documentsSubmitTransformTimeoutSeconds)
		pr.Done()
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

	documentsSubmitTransformCmd.Flags().StringVarP(&documentsSubmitTransformFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	documentsSubmitTransformCmd.MarkFlagFilename("file")
	documentsSubmitTransformCmd.Flags().StringVar(&documentsSubmitTransformArtifactSchema, "artifact-schema", "", "artifact schema")
	documentsSubmitTransformCmd.Flags().StringVar(&documentsSubmitTransformModel, "model", "", "model")
	documentsSubmitTransformCmd.Flags().StringVar(&documentsSubmitTransformPrompt, "prompt", "", "prompt")
	documentsSubmitTransformCmd.Flags().StringVar(&documentsSubmitTransformPromptId, "prompt-id", "", "prompt id")
	documentsSubmitTransformCmd.Flags().StringVar(&documentsSubmitTransformTimeoutSeconds, "timeout-seconds", "", "timeout seconds")
	documentsSubmitTransformCmd.Flags().BoolVar(&documentsSubmitTransformTrace, "trace", false, "stream parsing trace after upload")
	documentsSubmitTransformCmd.Flags().BoolVar(&documentsSubmitTransformBrowser, "browser", false, "open trace in console")
}
