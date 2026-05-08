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
	fileUploadsUploadAndListContentFile string
	fileUploadsUploadAndListContentTrace bool
	fileUploadsUploadAndListContentBrowser bool
)

var fileUploadsUploadAndListContentCmd = &cobra.Command{
	Use:   "and-list-content",
	Short: "Upload Content (sync)",
	Long:  `Upload Content (sync)`,
	Example: "meibel datasources file-uploads and-list-content",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if fileUploadsUploadAndListContentFile == "" {
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
				Value(&fileUploadsUploadAndListContentFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if fileUploadsUploadAndListContentFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(fileUploadsUploadAndListContentFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(fileUploadsUploadAndListContentFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		result, err := client.Datasources.FileUploads.UploadAndListContent(ctx, pr, fileName)
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

		if fileUploadsUploadAndListContentBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if fileUploadsUploadAndListContentTrace && jr.JobID != "" {
			output.Print(result)

			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
			defer cancel()

			stream, err := client.Datasources.FileUploads.StreamUploadProgress(ctx, jr.JobID)
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
	fileUploadsCmd.AddCommand(fileUploadsUploadAndListContentCmd)

	fileUploadsUploadAndListContentCmd.Flags().StringVarP(&fileUploadsUploadAndListContentFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	fileUploadsUploadAndListContentCmd.MarkFlagFilename("file")
	fileUploadsUploadAndListContentCmd.Flags().BoolVar(&fileUploadsUploadAndListContentTrace, "trace", false, "stream parsing trace after upload")
	fileUploadsUploadAndListContentCmd.Flags().BoolVar(&fileUploadsUploadAndListContentBrowser, "browser", false, "open trace in console")
}
