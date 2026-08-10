package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/config"
	"github.com/meibel-ai/meibel-cli/internal/tui"
	"github.com/meibel-ai/meibel-cli/internal/pathutil"
	sdk "github.com/meibel-ai/meibel-go/v2"
)

var (
	documentsSubmitDeepTransformFile string
	documentsSubmitDeepTransformSchema string
	documentsSubmitDeepTransformRootName string
	documentsSubmitDeepTransformGuidance string
	documentsSubmitDeepTransformMaxPages int64
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

		if documentsSubmitDeepTransformFile == "" && term.IsTerminal(int(os.Stdin.Fd())) {
			if err := huh.NewForm(huh.NewGroup(
				huh.NewInput().
					Title("File path").
					Description("Paste or type a path — leave blank to browse").
					Value(&documentsSubmitDeepTransformFile),
			)).Run(); err != nil {
				return err
			}
			documentsSubmitDeepTransformFile = pathutil.Expand(documentsSubmitDeepTransformFile)
		}

		if documentsSubmitDeepTransformFile == "" {
			picker := huh.NewFilePicker().
				Title("Select a file").
				CurrentDirectory(pathutil.StartDir()).
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

		opts := sdk.DocumentsSubmitDeepTransformOptions{}
		opts.File = documentsSubmitDeepTransformFile
		opts.Schema = documentsSubmitDeepTransformSchema
		if documentsSubmitDeepTransformRootName != "" {
			opts.RootName = &documentsSubmitDeepTransformRootName
		}
		if documentsSubmitDeepTransformGuidance != "" {
			opts.Guidance = &documentsSubmitDeepTransformGuidance
		}
		if documentsSubmitDeepTransformMaxPages != 0 {
			opts.MaxPages = &documentsSubmitDeepTransformMaxPages
		}

		if documentsSubmitDeepTransformWait {
			sp := tui.StartSpinner("Submit a deep-transform extraction from a file upload (async)")
			result, err := client.Documents.SubmitTransform(ctx, sdk.DocumentsSubmitTransformOptions{File: documentsSubmitDeepTransformFile, Schema: documentsSubmitDeepTransformSchema})
			sp.Stop()
			if err != nil {
				return err
			}
			return output.Print(result)
		}

		sp := tui.StartSpinner("Submit a deep-transform extraction from a file upload (async)")
		result, err := client.Documents.SubmitDeepTransform(ctx, opts)
		sp.Stop()
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

	documentsSubmitDeepTransformCmd.Flags().StringVarP(&documentsSubmitDeepTransformFile, "file", "f", "", "Document file to extract from")
	documentsSubmitDeepTransformCmd.MarkFlagFilename("file")
	documentsSubmitDeepTransformCmd.Flags().StringVar(&documentsSubmitDeepTransformSchema, "schema", "", "JSON Schema (as a JSON string) of the entities to extract")
	documentsSubmitDeepTransformCmd.MarkFlagRequired("schema")
	documentsSubmitDeepTransformCmd.Flags().StringVar(&documentsSubmitDeepTransformRootName, "root-name", "", "Name of the root entity in the schema. Optional: resolved from the schema's title or inferred when omitted.")
	documentsSubmitDeepTransformCmd.Flags().StringVar(&documentsSubmitDeepTransformGuidance, "guidance", "", "Optional domain guidance for the extraction")
	documentsSubmitDeepTransformCmd.Flags().Int64Var(&documentsSubmitDeepTransformMaxPages, "max-pages", 0, "Optional cap on the number of pages to process")
	documentsSubmitDeepTransformCmd.Flags().BoolVar(&documentsSubmitDeepTransformTrace, "trace", false, "stream parsing trace after upload")
	documentsSubmitDeepTransformCmd.Flags().BoolVar(&documentsSubmitDeepTransformBrowser, "browser", false, "open trace in console")
	documentsSubmitDeepTransformCmd.Flags().BoolVar(&documentsSubmitDeepTransformWait, "wait", false, "wait for completion (synchronous)")
}
