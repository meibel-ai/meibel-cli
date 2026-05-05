package tui

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Color palette
var (
	colorPrimary = lipgloss.Color("#7C3AED")
	colorSuccess = lipgloss.Color("#10B981")
	colorError   = lipgloss.Color("#EF4444")
	colorMuted   = lipgloss.Color("#6B7280")
)

// EventStream is a generic interface for streaming events
type EventStream[T any] interface {
	Events() <-chan T
	Errors() <-chan error
	Close() error
}

// StreamEvents displays streaming events in a TUI.
func StreamEvents[T any](ctx context.Context, stream EventStream[T]) error {
	p := tea.NewProgram(newStreamModel(ctx, stream))
	_, err := p.Run()
	return err
}

// streamModel is the Bubbletea model for streaming events.
type streamModel[T any] struct {
	ctx       context.Context
	stream    EventStream[T]
	events    []string
	spinner   spinner.Model
	err       error
	done      bool
	maxEvents int
}

type eventMsg string
type errorMsg error
type doneMsg struct{}

func newStreamModel[T any](ctx context.Context, stream EventStream[T]) *streamModel[T] {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(colorPrimary)

	return &streamModel[T]{
		ctx:       ctx,
		stream:    stream,
		events:    make([]string, 0),
		spinner:   s,
		maxEvents: 20, // Keep last N events
	}
}

func (m *streamModel[T]) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.listenForEvents(),
	)
}

func (m *streamModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.stream.Close()
			return m, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case eventMsg:
		m.events = append(m.events, string(msg))
		// Keep only last N events
		if len(m.events) > m.maxEvents {
			m.events = m.events[len(m.events)-m.maxEvents:]
		}
		return m, m.listenForEvents()

	case errorMsg:
		m.err = msg
		m.done = true
		return m, tea.Quit

	case doneMsg:
		m.done = true
		return m, tea.Quit
	}

	return m, nil
}

func (m *streamModel[T]) View() string {
	var b strings.Builder

	// Header
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(colorPrimary).
		MarginBottom(1)

	b.WriteString(headerStyle.Render("Streaming Events"))
	b.WriteString("\n\n")

	// Events
	eventStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F9FAFB"))

	mutedStyle := lipgloss.NewStyle().
		Foreground(colorMuted)

	if len(m.events) == 0 {
		b.WriteString(mutedStyle.Render("Waiting for events..."))
		b.WriteString("\n")
	} else {
		for _, event := range m.events {
			b.WriteString(eventStyle.Render("• " + event))
			b.WriteString("\n")
		}
	}

	b.WriteString("\n")

	// Status
	if m.err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(colorError).
			Bold(true)
		b.WriteString(errorStyle.Render("Error: " + m.err.Error()))
	} else if m.done {
		successStyle := lipgloss.NewStyle().
			Foreground(colorSuccess)
		b.WriteString(successStyle.Render("✓ Stream ended"))
	} else {
		b.WriteString(m.spinner.View())
		b.WriteString(mutedStyle.Render(" Receiving... (press q to quit)"))
	}

	return b.String()
}

func (m *streamModel[T]) listenForEvents() tea.Cmd {
	return func() tea.Msg {
		select {
		case <-m.ctx.Done():
			return doneMsg{}
		case err, ok := <-m.stream.Errors():
			if !ok {
				return doneMsg{}
			}
			return errorMsg(err)
		case event, ok := <-m.stream.Events():
			if !ok {
				return doneMsg{}
			}
			// Convert event to string for display
			data, err := json.Marshal(event)
			if err != nil {
				return eventMsg(fmt.Sprintf("%v", event))
			}
			return eventMsg(string(data))
		}
	}
}

// Spinner shows a spinner while an operation is in progress.
type Spinner struct {
	model   spinner.Model
	message string
}

// NewSpinner creates a new spinner with a message.
func NewSpinner(message string) *Spinner {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(colorPrimary)
	return &Spinner{
		model:   s,
		message: message,
	}
}

// WithSpinner runs a function with a spinner displayed.
func WithSpinner[T any](message string, fn func() (T, error)) (T, error) {
	// For simplicity, just run the function directly
	// A full implementation would use Bubbletea
	fmt.Print(message + "...")
	result, err := fn()
	if err != nil {
		fmt.Println(" ✗")
		return result, err
	}
	fmt.Println(" ✓")
	return result, nil
}
