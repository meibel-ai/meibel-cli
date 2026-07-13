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
	documentsSubmitDeepTransformFile string
	documentsSubmitDeepTransformSchema string
	documentsSubmitDeepTransformRootName string
	documentsSubmitDeepTransformGuidance string
	documentsSubmitDeepTransformMaxPages string
	documentsSubmitDeepTransformTrace bool
	documentsSubmitDeepTransformBrowser bool
	documentsSubmitDeepTransformWait bool
)

var documentsSubmitDeepTransformCmd = &cobra.Command{
	Use:   "submit-deep-transform",
	Short: "Submit a deep-transform extraction from a file upload (async)",
	Long:  `Upload a document and submit an extraction against a JSON schema, returning immediately with a job id. To reuse an already-parsed document instead of uploading, use POST /documents/deep-transform/from-document. Poll status via GET /documents/deep-transform/{job_id} and download artifacts once it succeeds. Submission is idempotent on the (document, schema) pair.`,
	Example: "meibel documents submit-deep-transform",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if documentsSubmitDeepTransformFile == "" {
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
				Value(&documentsSubmitDeepTransformFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if documentsSubmitDeepTransformFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(documentsSubmitDeepTransformFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(documentsSubmitDeepTransformFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		if documentsSubmitDeepTransformWait {
			result, err := client.Documents.SubmitTransform(ctx, pr, fileName, documentsSubmitDeepTransformSchema, documentsSubmitDeepTransformRootName, documentsSubmitDeepTransformGuidance, documentsSubmitDeepTransformMaxPages)
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

			if documentsSubmitDeepTransformBrowser && jr.JobID != "" {
				consoleURL := deriveConsoleURL(config.GetString("base_url"))
				projectID := config.GetString("project_id")
				if consoleURL != "" && projectID != "" {
					url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
					openBrowser(url)
				}
			}

			return output.Print(result)
		}

		result, err := client.Documents.SubmitDeepTransform(ctx, pr, fileName, documentsSubmitDeepTransformSchema, documentsSubmitDeepTransformRootName, documentsSubmitDeepTransformGuidance, documentsSubmitDeepTransformMaxPages)
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

		if documentsSubmitDeepTransformBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if documentsSubmitDeepTransformTrace && jr.JobID != "" {
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
	documentsCmd.AddCommand(documentsSubmitDeepTransformCmd)

	documentsSubmitDeepTransformCmd.Flags().StringVarP(&documentsSubmitDeepTransformFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	documentsSubmitDeepTransformCmd.MarkFlagFilename("file")
	documentsSubmitDeepTransformCmd.Flags().StringVar(&documentsSubmitDeepTransformSchema, "schema", "", "schema")
	documentsSubmitDeepTransformCmd.Flags().StringVar(&documentsSubmitDeepTransformRootName, "root-name", "", "root name")
	documentsSubmitDeepTransformCmd.Flags().StringVar(&documentsSubmitDeepTransformGuidance, "guidance", "", "guidance")
	documentsSubmitDeepTransformCmd.Flags().StringVar(&documentsSubmitDeepTransformMaxPages, "max-pages", "", "max pages")
	documentsSubmitDeepTransformCmd.Flags().BoolVar(&documentsSubmitDeepTransformTrace, "trace", false, "stream parsing trace after upload")
	documentsSubmitDeepTransformCmd.Flags().BoolVar(&documentsSubmitDeepTransformBrowser, "browser", false, "open trace in console")
	documentsSubmitDeepTransformCmd.Flags().BoolVar(&documentsSubmitDeepTransformWait, "wait", false, "wait for parsing to complete (synchronous)")
}
