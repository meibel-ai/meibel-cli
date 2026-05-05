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
	"github.com/meibel-ai/meibel/internal/output"
	"github.com/meibel-ai/meibel/internal/config"
	"github.com/meibel-ai/meibel/internal/tui"
	"github.com/meibel-ai/meibel/internal/upload"
)

var (
	documentsParseDocumentFile string
	documentsParseDocumentTrace bool
	documentsParseDocumentBrowser bool
	documentsParseDocumentWait bool
)

var documentsParseDocumentCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a document (async)",
	Long:  `Upload a document for asynchronous parsing. Returns a job ID to track progress.`,
	Example: "meibel documents parse",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if documentsParseDocumentFile == "" {
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
				Value(&documentsParseDocumentFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if documentsParseDocumentFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(documentsParseDocumentFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(documentsParseDocumentFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		if documentsParseDocumentWait {
			result, err := client.Documents.ProcessDocument(ctx, pr, fileName)
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

			if documentsParseDocumentBrowser && jr.JobID != "" {
				consoleURL := deriveConsoleURL(config.GetString("base_url"))
				projectID := config.GetString("project_id")
				if consoleURL != "" && projectID != "" {
					url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
					openBrowser(url)
				}
			}

			if !output.PrintMarkdown(result, "result") {
				return output.Print(result)
			}
			return nil
		}

		result, err := client.Documents.ParseDocument(ctx, pr, fileName)
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

		if documentsParseDocumentBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if documentsParseDocumentTrace && jr.JobID != "" {
			output.Print(result)

			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
			defer cancel()

			stream, err := client.Documents.StreamDocumentTrace(ctx, jr.JobID)
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
	documentsCmd.AddCommand(documentsParseDocumentCmd)

	documentsParseDocumentCmd.Flags().StringVarP(&documentsParseDocumentFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	documentsParseDocumentCmd.MarkFlagFilename("file")
	documentsParseDocumentCmd.Flags().BoolVar(&documentsParseDocumentTrace, "trace", false, "stream parsing trace after upload")
	documentsParseDocumentCmd.Flags().BoolVar(&documentsParseDocumentBrowser, "browser", false, "open trace in console")
	documentsParseDocumentCmd.Flags().BoolVar(&documentsParseDocumentWait, "wait", false, "wait for parsing to complete (synchronous)")
}
