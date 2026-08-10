package output

import (
	"strings"
	"unicode/utf8"
)

// SelectColumns decides which columns of a collection to show by default.
//
// Dumping every field of a response model produces tables several times wider
// than any terminal, so the default view is narrowed. Preferred names, when
// supplied from config, win outright; otherwise a heuristic keeps the columns
// that identify and summarise a row and discards the rest. `--wide` bypasses
// this entirely, and `--json` still carries every field.
//
// Returns indices into headers, in display order.
func SelectColumns(headers []string, rows [][]string, preferred []string, avail int) []int {
	if len(headers) == 0 {
		return nil
	}

	// Explicit configuration: honour the given order, ignoring names the response
	// does not actually contain so a stale config cannot blank the output.
	if len(preferred) > 0 {
		var idx []int
		for _, want := range preferred {
			for i, h := range headers {
				if strings.EqualFold(h, want) {
					idx = append(idx, i)
					break
				}
			}
		}
		if len(idx) > 0 {
			return idx
		}
	}

	widths := columnWidths(headers, rows)

	var kept []columnCandidate
	for i, h := range headers {
		// Always-empty columns tell the reader nothing.
		if widths[i] == 0 || allEmpty(rows, i) {
			continue
		}
		// Structured values render as Go literals ([], map[...]) and are noise in
		// a table; they remain available via --json.
		if looksStructured(rows, i) {
			continue
		}
		// Width deliberately does not exclude a column here. A single long value
		// would otherwise drop the whole column — one 61-character datasource name
		// was enough to lose `name` for every row. Prose is demoted in the ranking
		// instead, and truncation handles the long values.
		kept = append(kept, columnCandidate{index: i, rank: rankColumn(h, widths[i])})
	}
	if len(kept) == 0 {
		// Everything was filtered out; fall back to the natural order so the user
		// still sees something.
		idx := make([]int, len(headers))
		for i := range headers {
			idx[i] = i
		}
		return idx
	}

	// Take columns in rank order until the width budget is spent, then restore
	// the original field order so the layout matches the model.
	sortByRank(kept)

	var chosen []int
	used := 0
	for _, c := range kept {
		// Charge the budget the width the column will actually occupy. Reserving
		// only the floor admits far too many columns, which then all get squeezed
		// to an unreadable ~18 characters.
		w := widths[c.index]
		if w > maxColWidth {
			w = maxColWidth
		}
		// An identifier is only useful in full; never shrink one to fit.
		if isIdentifier(headers[c.index], widths[c.index]) {
			w = widths[c.index]
		}
		gap := 0
		if len(chosen) > 0 {
			gap = gutter
		}
		if used+gap+w > avail && len(chosen) > 0 {
			continue
		}
		chosen = append(chosen, c.index)
		used += gap + w
	}
	if len(chosen) == 0 {
		chosen = []int{kept[0].index}
	}

	sortInts(chosen)
	return chosen
}

// rankColumn scores how much a column earns its place. Lower sorts first.
func rankColumn(header string, width int) int {
	h := strings.ToLower(header)
	switch {
	case h == "name":
		return 0
	case h == "display_name" || h == "displayname":
		return 1
	case isIdentifier(header, width):
		return 2
	case strings.Contains(h, "status") || strings.Contains(h, "state"):
		return 3
	case h == "type" || h == "kind":
		return 4
	case width <= 12:
		// Short scalars — counts, flags, enums — pack a lot of signal per column.
		return 5
	case width > freeTextWidth:
		// Prose (descriptions, prompts) is last: useful only if space remains.
		return 8
	case strings.Contains(h, "_at") || strings.Contains(h, "time") || strings.Contains(h, "date"):
		return 7
	default:
		return 6
	}
}

// isIdentifier reports whether a column holds opaque handles that lose all
// meaning when truncated.
func isIdentifier(header string, width int) bool {
	h := strings.ToLower(header)
	if h == "id" || strings.HasSuffix(h, "_id") || strings.HasSuffix(h, "id") || strings.Contains(h, "urn") {
		return width >= idWidthThreshold
	}
	return false
}

func columnWidths(headers []string, rows [][]string) []int {
	widths := make([]int, len(headers))
	for i := range headers {
		widths[i] = utf8.RuneCountInString(headers[i])
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
	return widths
}

func allEmpty(rows [][]string, col int) bool {
	for _, row := range rows {
		if col < len(row) {
			v := strings.TrimSpace(row[col])
			if v != "" && v != "-" && v != "<nil>" {
				return false
			}
		}
	}
	return true
}

func looksStructured(rows [][]string, col int) bool {
	for _, row := range rows {
		if col >= len(row) {
			continue
		}
		v := strings.TrimSpace(row[col])
		if v == "" {
			continue
		}
		if strings.HasPrefix(v, "[") || strings.HasPrefix(v, "map[") || strings.HasPrefix(v, "{") {
			return true
		}
	}
	return false
}

// columnCandidate pairs a column index with its display priority.
type columnCandidate struct {
	index int
	rank  int
}

// Small local sorts keep this file free of a sort import in generated output.
func sortByRank(c []columnCandidate) {
	for i := 1; i < len(c); i++ {
		for j := i; j > 0 && c[j].rank < c[j-1].rank; j-- {
			c[j], c[j-1] = c[j-1], c[j]
		}
	}
}

func sortInts(v []int) {
	for i := 1; i < len(v); i++ {
		for j := i; j > 0 && v[j] < v[j-1]; j-- {
			v[j], v[j-1] = v[j-1], v[j]
		}
	}
}
