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
	documentsParseFile string
	documentsParseTrace bool
	documentsParseBrowser bool
	documentsParseWait bool
)

var documentsParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a document (async)",
	Long:  `Upload a document for asynchronous parsing. Returns a job ID to track progress.`,
	Example: "meibel documents parse",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if documentsParseFile == "" {
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
				Value(&documentsParseFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if documentsParseFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(documentsParseFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(documentsParseFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		if documentsParseWait {
			result, err := client.Documents.Process(ctx, pr, fileName)
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

			if documentsParseBrowser && jr.JobID != "" {
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

		result, err := client.Documents.Parse(ctx, pr, fileName)
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

		if documentsParseBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if documentsParseTrace && jr.JobID != "" {
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
	documentsCmd.AddCommand(documentsParseCmd)

	documentsParseCmd.Flags().StringVarP(&documentsParseFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	documentsParseCmd.MarkFlagFilename("file")
	documentsParseCmd.Flags().BoolVar(&documentsParseTrace, "trace", false, "stream parsing trace after upload")
	documentsParseCmd.Flags().BoolVar(&documentsParseBrowser, "browser", false, "open trace in console")
	documentsParseCmd.Flags().BoolVar(&documentsParseWait, "wait", false, "wait for parsing to complete (synchronous)")
}
