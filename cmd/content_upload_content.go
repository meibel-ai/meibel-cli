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
	contentUploadContentFile string
	contentUploadContentTrace bool
	contentUploadContentBrowser bool
	contentUploadContentWait bool
)

var contentUploadContentCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload Content (async)",
	Long:  `Upload Content (async)`,
	Example: "meibel content upload",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if contentUploadContentFile == "" {
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
				Value(&contentUploadContentFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if contentUploadContentFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(contentUploadContentFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(contentUploadContentFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		if contentUploadContentWait {
			result, err := client.Content.UploadAndListContent(ctx, pr, fileName)
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

			if contentUploadContentBrowser && jr.JobID != "" {
				consoleURL := deriveConsoleURL(config.GetString("base_url"))
				projectID := config.GetString("project_id")
				if consoleURL != "" && projectID != "" {
					url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
					openBrowser(url)
				}
			}

			return output.Print(result)
		}

		result, err := client.Content.UploadContent(ctx, pr, fileName)
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

		if contentUploadContentBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if contentUploadContentTrace && jr.JobID != "" {
			output.Print(result)

			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
			defer cancel()

			stream, err := client.Content.StreamUploadProgress(ctx, jr.JobID)
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
	contentCmd.AddCommand(contentUploadContentCmd)

	contentUploadContentCmd.Flags().StringVarP(&contentUploadContentFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	contentUploadContentCmd.MarkFlagFilename("file")
	contentUploadContentCmd.Flags().BoolVar(&contentUploadContentTrace, "trace", false, "stream parsing trace after upload")
	contentUploadContentCmd.Flags().BoolVar(&contentUploadContentBrowser, "browser", false, "open trace in console")
	contentUploadContentCmd.Flags().BoolVar(&contentUploadContentWait, "wait", false, "wait for parsing to complete (synchronous)")
}
