package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// Spinner is an in-flight activity indicator shown while a request runs.
//
// It renders to stderr so that piping stdout stays clean, and it reads no input
// so it cannot swallow stdin from commands that accept a piped request body.
// When stderr is not a terminal every method is a no-op, which keeps CI logs and
// piped output free of escape codes.
type Spinner struct {
	prog *tea.Program
	done chan struct{}
}

// StartSpinner begins displaying a spinner with the given message. Always call
// Stop, ideally via defer, or the program will keep drawing.
func StartSpinner(message string) *Spinner {
	if !term.IsTerminal(int(os.Stderr.Fd())) {
		return &Spinner{}
	}

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(colorPrimary)

	p := tea.NewProgram(
		&spinnerModel{spinner: s, message: message},
		tea.WithOutput(os.Stderr),
		// No input: a spinner needs none, and consuming stdin would break
		// commands that read a request body from a pipe.
		tea.WithInput(nil),
		// Leave signal handling to the command, which installs its own handlers
		// for streaming operations.
		tea.WithoutSignalHandler(),
	)

	sp := &Spinner{prog: p, done: make(chan struct{})}
	go func() {
		defer close(sp.done)
		_, _ = p.Run()
	}()
	return sp
}

// Stop halts the spinner and clears the line it occupied. Safe to call on a
// no-op spinner and safe to call more than once.
func (s *Spinner) Stop() {
	if s == nil || s.prog == nil {
		return
	}
	s.prog.Quit()
	<-s.done
	s.prog = nil
	// Bubbletea leaves its final frame on screen; erase it so the spinner does
	// not linger above the command's real output.
	fmt.Fprint(os.Stderr, "\r\033[K")
}

type spinnerModel struct {
	spinner spinner.Model
	message string
}

func (m *spinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := msg.(spinner.TickMsg); ok {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *spinnerModel) View() string {
	return m.spinner.View() + lipgloss.NewStyle().Foreground(colorMuted).Render(" "+m.message)
}
