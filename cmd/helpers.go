package cmd

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

func openBrowser(urlStr string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", urlStr)
	case "linux":
		cmd = exec.Command("xdg-open", urlStr)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", urlStr)
	}
	if cmd != nil {
		cmd.Start()
	}
}

func deriveConsoleURL(baseURL string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	host := u.Hostname()
	if strings.HasPrefix(host, "api.") {
		host = "console." + strings.TrimPrefix(host, "api.")
	}
	return fmt.Sprintf("%s://%s", u.Scheme, host)
}
