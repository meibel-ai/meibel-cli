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
	contentUploadAndListContentFile string
	contentUploadAndListContentTrace bool
	contentUploadAndListContentBrowser bool
)

var contentUploadAndListContentCmd = &cobra.Command{
	Use:   "upload-and-list",
	Short: "Upload Content (sync)",
	Long:  `Upload Content (sync)`,
	Example: "meibel content upload-and-list",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if contentUploadAndListContentFile == "" {
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
				Value(&contentUploadAndListContentFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if contentUploadAndListContentFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(contentUploadAndListContentFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(contentUploadAndListContentFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

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

		if contentUploadAndListContentBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if contentUploadAndListContentTrace && jr.JobID != "" {
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
	contentCmd.AddCommand(contentUploadAndListContentCmd)

	contentUploadAndListContentCmd.Flags().StringVarP(&contentUploadAndListContentFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	contentUploadAndListContentCmd.MarkFlagFilename("file")
	contentUploadAndListContentCmd.Flags().BoolVar(&contentUploadAndListContentTrace, "trace", false, "stream parsing trace after upload")
	contentUploadAndListContentCmd.Flags().BoolVar(&contentUploadAndListContentBrowser, "browser", false, "open trace in console")
}
