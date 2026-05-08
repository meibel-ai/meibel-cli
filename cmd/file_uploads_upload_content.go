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
	fileUploadsUploadContentFile string
	fileUploadsUploadContentDatasourceId string
	fileUploadsUploadContentName string
	fileUploadsUploadContentDescription string
	fileUploadsUploadContentMetadataConfig string
	fileUploadsUploadContentTrace bool
	fileUploadsUploadContentBrowser bool
	fileUploadsUploadContentWait bool
)

var fileUploadsUploadContentCmd = &cobra.Command{
	Use:   "content",
	Short: "Upload Content (async)",
	Long:  `Upload Content (async)`,
	Example: "meibel datasources file-uploads content",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if fileUploadsUploadContentFile == "" {
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
				Value(&fileUploadsUploadContentFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if fileUploadsUploadContentFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(fileUploadsUploadContentFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(fileUploadsUploadContentFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		if fileUploadsUploadContentWait {
			result, err := client.Datasources.FileUploads.UploadAndListContent(ctx, pr, fileName, fileUploadsUploadContentDatasourceId, fileUploadsUploadContentName, fileUploadsUploadContentDescription, fileUploadsUploadContentMetadataConfig)
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

			if fileUploadsUploadContentBrowser && jr.JobID != "" {
				consoleURL := deriveConsoleURL(config.GetString("base_url"))
				projectID := config.GetString("project_id")
				if consoleURL != "" && projectID != "" {
					url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
					openBrowser(url)
				}
			}

			return output.Print(result)
		}

		result, err := client.Datasources.FileUploads.UploadContent(ctx, pr, fileName, fileUploadsUploadContentDatasourceId, fileUploadsUploadContentName, fileUploadsUploadContentDescription, fileUploadsUploadContentMetadataConfig)
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

		if fileUploadsUploadContentBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if fileUploadsUploadContentTrace && jr.JobID != "" {
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
	fileUploadsCmd.AddCommand(fileUploadsUploadContentCmd)

	fileUploadsUploadContentCmd.Flags().StringVarP(&fileUploadsUploadContentFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	fileUploadsUploadContentCmd.MarkFlagFilename("file")
	fileUploadsUploadContentCmd.Flags().StringVar(&fileUploadsUploadContentDatasourceId, "datasource-id", "", "datasource id")
	fileUploadsUploadContentCmd.Flags().StringVar(&fileUploadsUploadContentName, "name", "", "name")
	fileUploadsUploadContentCmd.Flags().StringVar(&fileUploadsUploadContentDescription, "description", "", "description")
	fileUploadsUploadContentCmd.Flags().StringVar(&fileUploadsUploadContentMetadataConfig, "metadata-config", "", "metadata config")
	fileUploadsUploadContentCmd.Flags().BoolVar(&fileUploadsUploadContentTrace, "trace", false, "stream parsing trace after upload")
	fileUploadsUploadContentCmd.Flags().BoolVar(&fileUploadsUploadContentBrowser, "browser", false, "open trace in console")
	fileUploadsUploadContentCmd.Flags().BoolVar(&fileUploadsUploadContentWait, "wait", false, "wait for parsing to complete (synchronous)")
}
