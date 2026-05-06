module github.com/meibel-ai/meibel

go 1.21

require (
	github.com/meibel-ai/meibel-go/v2 v2.0.0
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.2
	github.com/charmbracelet/bubbletea v0.25.0
	github.com/charmbracelet/lipgloss v0.9.1
	github.com/charmbracelet/huh v1.0.0
	github.com/charmbracelet/bubbles v0.17.1
	github.com/charmbracelet/glamour v0.8.0
	golang.org/x/term v0.16.0
)

replace github.com/meibel-ai/meibel-go/v2 => ../go
