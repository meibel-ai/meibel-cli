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
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/config"
	"github.com/meibel-ai/meibel-cli/internal/tui"
	"github.com/meibel-ai/meibel-cli/internal/upload"
)

var (
	fileUploadUploadAndListContentFile string
	fileUploadUploadAndListContentTrace bool
	fileUploadUploadAndListContentBrowser bool
)

var fileUploadUploadAndListContentCmd = &cobra.Command{
	Use:   "and-list-content",
	Short: "Upload Content (sync)",
	Long:  `Upload Content (sync)`,
	Example: "meibel datasources file-upload and-list-content",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if fileUploadUploadAndListContentFile == "" {
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
				Value(&fileUploadUploadAndListContentFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if fileUploadUploadAndListContentFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(fileUploadUploadAndListContentFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(fileUploadUploadAndListContentFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		result, err := client.Datasources.FileUpload.UploadAndListContent(ctx, pr, fileName)
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

		if fileUploadUploadAndListContentBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if fileUploadUploadAndListContentTrace && jr.JobID != "" {
			output.Print(result)

			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
			defer cancel()

			stream, err := client.Datasources.FileUpload.StreamUploadProgress(ctx, jr.JobID)
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
	fileUploadCmd.AddCommand(fileUploadUploadAndListContentCmd)

	fileUploadUploadAndListContentCmd.Flags().StringVarP(&fileUploadUploadAndListContentFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	fileUploadUploadAndListContentCmd.MarkFlagFilename("file")
	fileUploadUploadAndListContentCmd.Flags().BoolVar(&fileUploadUploadAndListContentTrace, "trace", false, "stream parsing trace after upload")
	fileUploadUploadAndListContentCmd.Flags().BoolVar(&fileUploadUploadAndListContentBrowser, "browser", false, "open trace in console")
}
