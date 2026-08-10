package output

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// truncate shortens s to width runes, ending in an ellipsis when it had to cut.
func truncate(s string, width int) string {
	if utf8.RuneCountInString(s) <= width {
		return s
	}
	if width <= 1 {
		return strings.Repeat("…", width)
	}
	r := []rune(s)
	return string(r[:width-1]) + "…"
}

// maxColWidth caps any single column so one long free-text field (a description,
// a prompt) cannot push the table past the terminal and wrap every row.
const maxColWidth = 40

// PrintTable prints data as a formatted table.
func PrintTable(headers []string, rows [][]string) {
	if len(headers) == 0 {
		return
	}

	// Calculate column widths
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = utf8.RuneCountInString(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				if n := utf8.RuneCountInString(cell); n > widths[i] {
					widths[i] = n
				}
			}
		}
	}
	for i, w := range widths {
		if w > maxColWidth {
			widths[i] = maxColWidth
		}
	}

	// Truncate cells that exceed the capped width
	for _, row := range rows {
		for i := range row {
			if i < len(widths) {
				row[i] = truncate(row[i], widths[i])
			}
		}
	}
	for i := range headers {
		headers[i] = truncate(headers[i], widths[i])
	}

	// Print header
	var headerCells []string
	for i, h := range headers {
		headerCells = append(headerCells, Styles.TableHeader.Render(padRight(h, widths[i])))
	}
	fmt.Println(strings.Join(headerCells, "  "))

	// Print separator
	var sepCells []string
	for _, w := range widths {
		sepCells = append(sepCells, Styles.TableBorder.Render(strings.Repeat("─", w)))
	}
	fmt.Println(strings.Join(sepCells, "  "))

	// Print rows
	for _, row := range rows {
		var cells []string
		for i, cell := range row {
			if i < len(widths) {
				cells = append(cells, Styles.TableCell.Render(padRight(cell, widths[i])))
			}
		}
		fmt.Println(strings.Join(cells, ""))
	}
}

// PrintKeyValue prints a list of key-value pairs.
func PrintKeyValue(pairs map[string]string) {
	// Find max key width
	maxWidth := 0
	for k := range pairs {
		if len(k) > maxWidth {
			maxWidth = len(k)
		}
	}

	for k, v := range pairs {
		fmt.Printf("%s  %s\n",
			Styles.Bold.Render(padRight(k+":", maxWidth+1)),
			v)
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
		fmt.Printf("%s %s\n",
			Styles.Muted.Render(fmt.Sprintf("%d.", i+1)),
			item)
	}
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
