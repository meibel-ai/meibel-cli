package pathutil

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExpandPastedForms(t *testing.T) {
	home, _ := os.UserHomeDir()
	// A path with spaces and parentheses — the shapes that arrive escaped.
	real := "/tmp/Demo Files/Report PDS (624746).pdf"

	cases := []struct{ name, in, want string }{
		{"plain", real, real},
		{"double quoted", `"` + real + `"`, real},
		{"single quoted", `'` + real + `'`, real},
		{"drag-and-drop escapes", `/tmp/Demo\ Files/Report\ PDS\ \(624746\).pdf`, real},
		{"trailing whitespace", real + "  \n", real},
		{"tilde", "~/Documents/x.pdf", filepath.Join(home, "Documents/x.pdf")},
		{"quoted tilde", `"~/Documents/x.pdf"`, filepath.Join(home, "Documents/x.pdf")},
		{"empty", "   ", ""},
	}
	for _, c := range cases {
		if got := Expand(c.in); got != c.want {
			t.Errorf("%s:\n  in   %q\n  got  %q\n  want %q", c.name, c.in, got, c.want)
		}
	}
}

func TestStartDirIsWorkingDir(t *testing.T) {
	wd, _ := os.Getwd()
	if got := StartDir(); got != wd {
		t.Errorf("StartDir() = %q, want working dir %q", got, wd)
	}
}
