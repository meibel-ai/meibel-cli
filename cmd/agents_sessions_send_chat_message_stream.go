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
	"github.com/meibel-ai/meibel-go/meibel/internal/config"
	"github.com/meibel-ai/meibel-go/meibel/internal/tui"
	"github.com/meibel-ai/meibel-go/meibel/internal/upload"
)

var (
	agentsSessionsSendChatMessageStreamFile string
	agentsSessionsSendChatMessageStreamUserMessage string
	agentsSessionsSendChatMessageStreamTimeoutSeconds string
	agentsSessionsSendChatMessageStreamIncludeThinking string
	agentsSessionsSendChatMessageStreamIncludeToolActivity string
	agentsSessionsSendChatMessageStreamFiles string
	agentsSessionsSendChatMessageStreamTrace bool
	agentsSessionsSendChatMessageStreamBrowser bool
)

var agentsSessionsSendChatMessageStreamCmd = &cobra.Command{
	Use:   "send-chat-message-stream <session-id>",
	Short: "Send a chat message with file attachments and stream the response via SSE",
	Long:  `Send a chat message with file attachments and stream the response via SSE

Arguments:
  session-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel agents agents-sessions send-chat-message-stream <session-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		sessionId := args[0]

		if agentsSessionsSendChatMessageStreamFile == "" {
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
				Value(&agentsSessionsSendChatMessageStreamFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if agentsSessionsSendChatMessageStreamFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(agentsSessionsSendChatMessageStreamFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("failed to stat file: %w", err)
		}
		fileName := filepath.Base(agentsSessionsSendChatMessageStreamFile)
		pr := upload.NewProgressReader(f, fi.Size(), "Uploading")

		result, err := client.Agents.Sessions.SendChatMessageStream(ctx, sessionId, pr, fileName, agentsSessionsSendChatMessageStreamUserMessage, agentsSessionsSendChatMessageStreamTimeoutSeconds, agentsSessionsSendChatMessageStreamIncludeThinking, agentsSessionsSendChatMessageStreamIncludeToolActivity, agentsSessionsSendChatMessageStreamFiles)
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

		if agentsSessionsSendChatMessageStreamBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if agentsSessionsSendChatMessageStreamTrace && jr.JobID != "" {
			output.Print(result)

			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
			defer cancel()

			stream, err := client.Agents.Sessions.SendChatMessageStream(ctx, jr.JobID)
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
	agentsSessionsCmd.AddCommand(agentsSessionsSendChatMessageStreamCmd)

	agentsSessionsSendChatMessageStreamCmd.Flags().StringVarP(&agentsSessionsSendChatMessageStreamFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	agentsSessionsSendChatMessageStreamCmd.MarkFlagFilename("file")
	agentsSessionsSendChatMessageStreamCmd.Flags().StringVar(&agentsSessionsSendChatMessageStreamUserMessage, "user-message", "", "user message")
	agentsSessionsSendChatMessageStreamCmd.Flags().StringVar(&agentsSessionsSendChatMessageStreamTimeoutSeconds, "timeout-seconds", "", "timeout seconds")
	agentsSessionsSendChatMessageStreamCmd.Flags().StringVar(&agentsSessionsSendChatMessageStreamIncludeThinking, "include-thinking", "", "include thinking")
	agentsSessionsSendChatMessageStreamCmd.Flags().StringVar(&agentsSessionsSendChatMessageStreamIncludeToolActivity, "include-tool-activity", "", "include tool activity")
	agentsSessionsSendChatMessageStreamCmd.Flags().StringVar(&agentsSessionsSendChatMessageStreamFiles, "files", "", "files")
	agentsSessionsSendChatMessageStreamCmd.Flags().BoolVar(&agentsSessionsSendChatMessageStreamTrace, "trace", false, "stream parsing trace after upload")
	agentsSessionsSendChatMessageStreamCmd.Flags().BoolVar(&agentsSessionsSendChatMessageStreamBrowser, "browser", false, "open trace in console")
}
