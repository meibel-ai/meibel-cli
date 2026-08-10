package output

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"
)

// Layout constants for table rendering.
const (
	// gutter is the space between columns.
	gutter = 2
	// minColWidth is the narrowest a column may be squeezed before it is dropped
	// instead — narrower than this and the content is unreadable anyway.
	minColWidth = 6
	// fallbackWidth is used when the terminal size cannot be determined.
	fallbackWidth = 100
	// maxColWidth is the widest a single non-identifier column may claim, so one
	// long field cannot crowd out every other column.
	maxColWidth = 40
	// idWidthThreshold: values at least this wide that look like identifiers are
	// never truncated, because half a UUID cannot be copied or passed to another
	// command. Such a column is shown in full or dropped entirely.
	idWidthThreshold = 20
	// freeTextWidth: columns whose widest value exceeds this are treated as prose
	// (descriptions, prompts) and dropped from the default view.
	freeTextWidth = 60
)

// TerminalWidth returns the usable width for output.
//
// COLUMNS wins when set, which lets output be laid out for a width other than
// the current terminal's — useful when piping into a pager that side-scrolls, and
// the only way to get a deterministic width in tests. Falls back to a sane
// default when the size cannot be determined at all.
func TerminalWidth() int {
	if v := os.Getenv("COLUMNS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			return n
		}
	}
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || w <= 0 {
		return fallbackWidth
	}
	return w
}

// PrintTable prints headers and rows, fitted to the terminal width.
func PrintTable(headers []string, rows [][]string) {
	if len(headers) == 0 {
		return
	}
	widths := fitWidths(headers, rows, TerminalWidth())

	headerCells := make([]string, 0, len(headers))
	sepCells := make([]string, 0, len(headers))
	for i, h := range headers {
		headerCells = append(headerCells, Styles.TableHeader.Render(padRight(truncate(strings.ToUpper(h), widths[i]), widths[i])))
		sepCells = append(sepCells, Styles.TableBorder.Render(strings.Repeat("─", widths[i])))
	}
	sep := strings.Repeat(" ", gutter)
	fmt.Println(strings.Join(headerCells, sep))
	fmt.Println(strings.Join(sepCells, sep))

	for _, row := range rows {
		cells := make([]string, 0, len(widths))
		for i := range widths {
			v := ""
			if i < len(row) {
				v = row[i]
			}
			cells = append(cells, padRight(truncate(v, widths[i]), widths[i]))
		}
		fmt.Println(strings.Join(cells, sep))
	}
}

// fitWidths sizes columns to the natural content width, then squeezes the widest
// flexible column repeatedly until the whole table fits.
func fitWidths(headers []string, rows [][]string, avail int) []int {
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = utf8.RuneCountInString(h)
	}
	for _, row := range rows {
		for i := range widths {
			if i < len(row) {
				if n := utf8.RuneCountInString(row[i]); n > widths[i] {
					widths[i] = n
				}
			}
		}
	}

	budget := avail - gutter*(len(widths)-1)
	if budget < len(widths)*minColWidth {
		// Too many columns for this terminal to show anything useful; give each
		// an equal sliver rather than wrapping every row.
		for i := range widths {
			widths[i] = max(minColWidth, budget/len(widths))
		}
		return widths
	}

	for total(widths) > budget {
		// Squeeze the widest column that is still above the floor.
		victim, best := -1, minColWidth
		for i, w := range widths {
			if w > best {
				victim, best = i, w
			}
		}
		if victim < 0 {
			break
		}
		widths[victim]--
	}
	return widths
}

func total(widths []int) int {
	sum := 0
	for _, w := range widths {
		sum += w
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// PrintKeyValue prints a list of key-value pairs.
func PrintKeyValue(pairs map[string]string) {
	maxWidth := 0
	for k := range pairs {
		if n := utf8.RuneCountInString(k); n > maxWidth {
			maxWidth = n
		}
	}
	for k, v := range pairs {
		fmt.Printf("%s  %s\n", Styles.Bold.Render(padRight(k+":", maxWidth+1)), v)
	}
}

// PrintList prints a bulleted list.
func PrintList(items []string) {
	for _, item := range items {
		fmt.Printf("%s %s\n", Styles.Muted.Render(IconBullet), item)
	}
}

// PrintNumberedList prints a numbered list.
func PrintNumberedList(items []string) {
	for i, item := range items {
		fmt.Printf("%s %s\n", Styles.Muted.Render(fmt.Sprintf("%d.", i+1)), item)
	}
}

// truncate shortens s to width runes, ending in an ellipsis when it had to cut.
func truncate(s string, width int) string {
	if width <= 0 {
		return ""
	}
	if utf8.RuneCountInString(s) <= width {
		return s
	}
	if width == 1 {
		return "…"
	}
	return string([]rune(s)[:width-1]) + "…"
}

func padRight(s string, width int) string {
	// Rune count, not byte length: a multi-byte cell would otherwise be padded
	// short and knock the whole column out of alignment.
	n := utf8.RuneCountInString(s)
	if n >= width {
		return s
	}
	return s + strings.Repeat(" ", width-n)
}
