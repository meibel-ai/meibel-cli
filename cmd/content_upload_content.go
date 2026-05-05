package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"github.com/meibel-ai/meibel/internal/output"
	"github.com/meibel-ai/meibel/internal/upload"
)

var (
	contentUploadContentFile string
)

var contentUploadContentCmd = &cobra.Command{
	Use:   "upload <datasource-id>",
	Short: "Upload Content",
	Long:  `Upload Content

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources content upload <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

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

		result, err := client.Datasources.Content.UploadContent(ctx, datasourceId, pr, fileName)
		pr.Done()
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentUploadContentCmd)

	contentUploadContentCmd.Flags().StringVarP(&contentUploadContentFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	contentUploadContentCmd.MarkFlagFilename("file")
}
