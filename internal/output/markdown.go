package output

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/glamour"
)

// PrintMarkdown extracts a field from result and renders it as terminal markdown.
// If field is ".", treats the entire result as markdown when it's a string or *string.
// Returns false if not in a terminal, the field wasn't a string, or rendering failed
// (caller should fall back to Print).
func PrintMarkdown(result interface{}, field string) bool {
	if !IsTerminal() {
		return false
	}

	var md string
	if field == "." {
		switch v := result.(type) {
		case string:
			md = v
		case *string:
			if v == nil {
				return false
			}
			md = *v
		default:
			return false
		}
	} else {
		b, err := json.Marshal(result)
		if err != nil {
			return false
		}
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			return false
		}
		s, ok := m[field].(string)
		if !ok {
			return false
		}
		md = s
	}
	rendered, err := glamour.Render(md, "auto")
	if err != nil {
		return false
	}
	fmt.Print(rendered)
	return true
}
