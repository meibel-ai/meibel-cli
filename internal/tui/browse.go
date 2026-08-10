package tui

import (
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// BrowseTable opens a scrollable table for exploring a long result set.
//
// This is the opt-in counterpart to the static table: default output stays a
// plain width-fitted dump so it can be piped and kept in scrollback. Here the
// full column set is available by panning horizontally, which is the reason to
// reach for it — nothing is hidden to make the rows fit.
//
// Reports whether it ran; callers fall back to static output when it did not,
// which happens when stdout is not a terminal.
func BrowseTable(headers []string, rows [][]string) (bool, error) {
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return false, nil
	}
	if len(headers) == 0 {
		return false, nil
	}

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= 0 {
		width, height = 100, 24
	}

	m := &browseModel{
		headers: headers,
		rows:    rows,
		width:   width,
		height:  height,
	}
	m.rebuild()

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return false, err
	}
	return true, nil
}

type browseModel struct {
	headers []string
	rows    [][]string
	table   table.Model
	offset  int // index of the leftmost visible column
	width   int
	height  int
}

// rebuild recomputes the visible column window. bubbles/table has no horizontal
// scrolling, so panning is implemented by re-slicing the columns it is given.
func (m *browseModel) rebuild() {
	widths := make([]int, len(m.headers))
	for i, h := range m.headers {
		widths[i] = len([]rune(h))
	}
	for _, row := range m.rows {
		for i := range widths {
			if i < len(row) {
				if n := len([]rune(row[i])); n > widths[i] {
					widths[i] = n
				}
			}
		}
	}
	for i := range widths {
		if widths[i] > 40 {
			widths[i] = 40
		}
	}

	// Fit as many columns as the terminal allows, starting from the offset.
	var cols []table.Column
	var idx []int
	used := 0
	for i := m.offset; i < len(m.headers); i++ {
		w := widths[i]
		if used+w+2 > m.width-4 && len(cols) > 0 {
			break
		}
		cols = append(cols, table.Column{Title: strings.ToUpper(m.headers[i]), Width: w})
		idx = append(idx, i)
		used += w + 2
	}

	trows := make([]table.Row, 0, len(m.rows))
	for _, row := range m.rows {
		tr := make(table.Row, 0, len(idx))
		for _, i := range idx {
			v := ""
			if i < len(row) {
				v = row[i]
			}
			tr = append(tr, v)
		}
		trows = append(trows, tr)
	}

	h := m.height - 6
	if h < 3 {
		h = 3
	}

	t := table.New(
		table.WithColumns(cols),
		table.WithRows(trows),
		table.WithFocused(true),
		table.WithHeight(h),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#6B7280")).
		BorderBottom(true).
		Bold(true).
		Foreground(lipgloss.Color("#7C3AED"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#F9FAFB")).
		Background(lipgloss.Color("#7C3AED")).
		Bold(true)
	t.SetStyles(s)

	// Keep the cursor where it was across a re-slice so panning does not jump.
	if m.table.Cursor() > 0 && m.table.Cursor() < len(trows) {
		t.SetCursor(m.table.Cursor())
	}
	m.table = t
}

func (m *browseModel) Init() tea.Cmd { return nil }

func (m *browseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.rebuild()
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "right", "l":
			if m.offset < len(m.headers)-1 {
				m.offset++
				m.rebuild()
			}
			return m, nil
		case "left", "h":
			if m.offset > 0 {
				m.offset--
				m.rebuild()
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *browseModel) View() string {
	help := lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280")).Render(
		"↑/↓ row · ←/→ columns · q quit")
	pos := lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280")).Render(
		columnRange(m.offset, len(m.table.Columns()), len(m.headers)))
	return m.table.View() + "\n" + help + "  " + pos + "\n"
}

func columnRange(offset, shown, total int) string {
	if shown >= total {
		return ""
	}
	last := offset + shown
	return "cols " + itoa(offset+1) + "-" + itoa(last) + " of " + itoa(total)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
