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
	documentsTransformFile string
	documentsTransformArtifactSchema string
	documentsTransformModel string
	documentsTransformPrompt string
	documentsTransformPromptId string
	documentsTransformTimeoutSeconds string
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

		if documentsTransformFile == "" {
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
				Value(&documentsTransformFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if documentsTransformFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(documentsTransformFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(documentsTransformFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		if documentsTransformWait {
			result, err := client.Documents.SubmitTransform(ctx, pr, fileName, documentsTransformArtifactSchema, documentsTransformModel, documentsTransformPrompt, documentsTransformPromptId, documentsTransformTimeoutSeconds)
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

			if documentsTransformBrowser && jr.JobID != "" {
				consoleURL := deriveConsoleURL(config.GetString("base_url"))
				projectID := config.GetString("project_id")
				if consoleURL != "" && projectID != "" {
					url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
					openBrowser(url)
				}
			}

			return output.Print(result)
		}

		result, err := client.Documents.Transform(ctx, pr, fileName, documentsTransformArtifactSchema, documentsTransformModel, documentsTransformPrompt, documentsTransformPromptId, documentsTransformTimeoutSeconds)
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

	documentsTransformCmd.Flags().StringVarP(&documentsTransformFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	documentsTransformCmd.MarkFlagFilename("file")
	documentsTransformCmd.Flags().StringVar(&documentsTransformArtifactSchema, "artifact-schema", "", "artifact schema")
	documentsTransformCmd.Flags().StringVar(&documentsTransformModel, "model", "", "model")
	documentsTransformCmd.Flags().StringVar(&documentsTransformPrompt, "prompt", "", "prompt")
	documentsTransformCmd.Flags().StringVar(&documentsTransformPromptId, "prompt-id", "", "prompt id")
	documentsTransformCmd.Flags().StringVar(&documentsTransformTimeoutSeconds, "timeout-seconds", "", "timeout seconds")
	documentsTransformCmd.Flags().BoolVar(&documentsTransformTrace, "trace", false, "stream parsing trace after upload")
	documentsTransformCmd.Flags().BoolVar(&documentsTransformBrowser, "browser", false, "open trace in console")
	documentsTransformCmd.Flags().BoolVar(&documentsTransformWait, "wait", false, "wait for parsing to complete (synchronous)")
}
