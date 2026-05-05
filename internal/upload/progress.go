package upload

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	barFilled  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7C3AED"))
	barEmpty   = lipgloss.NewStyle().Foreground(lipgloss.Color("#374151"))
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9CA3AF"))
	pctStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#F9FAFB")).Bold(true)
)

// ProgressReader wraps an io.Reader and prints upload progress to stderr.
type ProgressReader struct {
	reader  io.Reader
	total   int64
	current int64
	label   string
	lastPct int
}

// NewProgressReader creates a reader that displays a progress bar as bytes are read.
func NewProgressReader(reader io.Reader, total int64, label string) *ProgressReader {
	return &ProgressReader{
		reader:  reader,
		total:   total,
		label:   label,
		lastPct: -1,
	}
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	if n > 0 {
		pr.current += int64(n)
		pct := int(float64(pr.current) / float64(pr.total) * 100)
		if pct > 100 {
			pct = 100
		}
		if pct != pr.lastPct {
			pr.lastPct = pct
			pr.render(pct)
		}
	}
	return n, err
}

func (pr *ProgressReader) render(pct int) {
	const width = 30
	filled := width * pct / 100
	empty := width - filled

	bar := barFilled.Render(strings.Repeat("█", filled)) +
		barEmpty.Render(strings.Repeat("░", empty))

	line := fmt.Sprintf("\r  %s %s %s",
		labelStyle.Render(pr.label),
		bar,
		pctStyle.Render(fmt.Sprintf("%3d%%", pct)),
	)
	fmt.Fprint(os.Stderr, line)
}

// Done prints a newline after the progress bar completes.
func (pr *ProgressReader) Done() {
	fmt.Fprintln(os.Stderr)
}
