package pathutil

import (
	"os"
	"path/filepath"
	"strings"
)

// Expand normalises a path a user pasted or typed into one os.Open accepts.
//
// Pasted paths arrive decorated in ways a shell would have stripped: quotes from
// copying a quoted argument, backslash-escaped spaces from dragging a file into a
// terminal, and a leading ~ that only a shell expands. Handling them here means a
// pasted path works as often as a shell-completed one.
func Expand(p string) string {
	p = strings.TrimSpace(p)
	if p == "" {
		return ""
	}

	// Surrounding quotes, as produced by copying a quoted shell argument.
	if len(p) >= 2 {
		if (p[0] == '\'' && p[len(p)-1] == '\'') || (p[0] == '"' && p[len(p)-1] == '"') {
			p = p[1 : len(p)-1]
		}
	}

	// Shell escapes from a drag-and-drop, e.g. "Demo\ Files". Only unescape
	// characters a shell would have escaped, so Windows-style separators and
	// literal backslashes in a name survive.
	if strings.Contains(p, `\`) {
		var b strings.Builder
		for i := 0; i < len(p); i++ {
			if p[i] == '\\' && i+1 < len(p) && strings.ContainsRune(" ()[]{}'\"&;$`*?!~|<>#", rune(p[i+1])) {
				continue
			}
			b.WriteByte(p[i])
		}
		p = b.String()
	}

	// Leading ~, which reaches the program unexpanded when it was quoted.
	if p == "~" || strings.HasPrefix(p, "~/") {
		if home, err := os.UserHomeDir(); err == nil {
			if p == "~" {
				p = home
			} else {
				p = filepath.Join(home, p[2:])
			}
		}
	}

	return p
}

// StartDir is where an interactive file browser should open: the working
// directory, since that is nearly always nearer the target than $HOME.
func StartDir() string {
	if wd, err := os.Getwd(); err == nil && wd != "" {
		return wd
	}
	if home, err := os.UserHomeDir(); err == nil && home != "" {
		return home
	}
	return "."
}
