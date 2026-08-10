package output

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"golang.org/x/term"
)

// Format represents the output format.
type Format int

const (
	// FormatAuto automatically selects format based on terminal
	FormatAuto Format = iota
	// FormatJSON outputs JSON
	FormatJSON
	// FormatTable outputs formatted tables
	FormatTable
	// FormatText outputs plain text
	FormatText
)

var (
	currentFormat = FormatAuto
	// wideOutput disables column narrowing (--wide).
	wideOutput bool
	// browseOutput requests the interactive table (--interactive).
	browseOutput bool
	// activePreferred carries the configured columns for the current Print call.
	activePreferred []string
)

// SetWide shows every column instead of the narrowed default view.
func SetWide(w bool) { wideOutput = w }

// SetBrowse requests the scrollable table for collection output.
func SetBrowse(b bool) { browseOutput = b }

// BrowseFunc renders a scrollable table. It is injected by the CLI to avoid an
// import cycle between output and tui, and is nil when browsing is unavailable.
var BrowseFunc func(headers []string, rows [][]string) (bool, error)

// SetFormat sets the global output format.
func SetFormat(f Format) {
	currentFormat = f
}

// GetFormat returns the current output format.
func GetFormat() Format {
	return currentFormat
}

// IsTerminal returns true if stdout is a terminal.
func IsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// Print outputs the data in the appropriate format.
//
// preferredColumns, when given, names the columns a collection should show in
// table view; it comes from config and is ignored for non-collections and for
// JSON output.
func Print(data interface{}, preferredColumns ...string) error {
	activePreferred = preferredColumns
	format := currentFormat

	// Auto-detect format
	if format == FormatAuto {
		if IsTerminal() {
			format = FormatTable
		} else {
			format = FormatJSON
		}
	}

	// A string result is already a document — markdown from a text endpoint, or
	// JSON text from the same endpoint under ?format=json. Emit it verbatim in
	// every mode: JSON-encoding it would escape the whole payload into one quoted
	// line, breaking `> out.md` and `| jq` alike.
	if s, ok := asString(data); ok {
		fmt.Print(s)
		if !strings.HasSuffix(s, "\n") {
			fmt.Println()
		}
		return nil
	}

	switch format {
	case FormatJSON:
		return printJSON(data)
	case FormatTable:
		return printTable(data)
	case FormatText:
		return printText(data)
	default:
		return printJSON(data)
	}
}

// PrintSuccess prints a success message.
func PrintSuccess(message string) {
	if IsTerminal() {
		fmt.Println(Styles.Success.Render(IconSuccess + " " + message))
	} else {
		fmt.Println(message)
	}
}

// PrintError prints an error message.
func PrintError(message string) {
	if IsTerminal() {
		fmt.Fprintln(os.Stderr, Styles.Error.Render(IconError+" "+message))
	} else {
		fmt.Fprintln(os.Stderr, "Error: "+message)
	}
}

// PrintWarning prints a warning message.
func PrintWarning(message string) {
	if IsTerminal() {
		fmt.Println(Styles.Warning.Render(IconWarning + " " + message))
	} else {
		fmt.Println("Warning: " + message)
	}
}

// PrintInfo prints an info message.
func PrintInfo(message string) {
	if IsTerminal() {
		fmt.Println(Styles.Muted.Render(IconInfo + " " + message))
	} else {
		fmt.Println(message)
	}
}

func printJSON(data interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func printText(data interface{}) error {
	fmt.Printf("%v\n", data)
	return nil
}

func printTable(data interface{}) error {
	// Handle nil
	if data == nil {
		fmt.Println("(no data)")
		return nil
	}

	// Use reflection to handle the data
	v := reflect.ValueOf(data)

	// Dereference pointer
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			fmt.Println("(no data)")
			return nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		if v.Len() == 0 {
			fmt.Println("(no items)")
			return nil
		}
		return printSliceAsTable(v)

	case reflect.Struct:
		// Most list endpoints wrap their collection in an envelope — {"data": [...],
		// "pagination": {...}} or {"datasources": [...]} — which as a struct would
		// render as one unreadable line of Go literals. Render the collection.
		if coll, ok := unwrapCollection(v); ok {
			return printSliceAsTable(coll)
		}
		return printStructAsTable(v)

	case reflect.Map:
		return printMapAsTable(v)

	default:
		// Fall back to JSON for unknown types
		return printJSON(data)
	}
}

// asString reports whether data is a string or *string, returning its value.
func asString(data interface{}) (string, bool) {
	switch v := data.(type) {
	case string:
		return v, true
	case *string:
		if v == nil {
			return "", false
		}
		return *v, true
	}
	return "", false
}

// jsonFieldName returns the wire name of a struct field, preferring its json tag.
// The tag is only sometimes suffixed with options, so the comma is optional —
// treating it as required left untagged-looking headers like "CreatedAt" beside
// tagged ones like "last_sync_at".
func jsonFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "" || tag == "-" {
		return field.Name
	}
	if i := strings.IndexByte(tag, ','); i >= 0 {
		tag = tag[:i]
	}
	if tag == "" {
		return field.Name
	}
	return tag
}

// unwrapCollection finds the single slice-of-structs field of an envelope struct,
// so a wrapped collection is tabulated rather than dumped as a literal. Structs
// with no such field, or more than one, are left alone — the intent would be
// ambiguous.
func unwrapCollection(v reflect.Value) (reflect.Value, bool) {
	t := v.Type()
	found := -1
	for i := 0; i < t.NumField(); i++ {
		if !t.Field(i).IsExported() {
			continue
		}
		f := unwrapValue(v.Field(i))
		if f.Kind() != reflect.Slice && f.Kind() != reflect.Array {
			continue
		}
		// Only element types that tabulate are worth unwrapping; a []string field
		// is data on the object, not the object's collection.
		elem := t.Field(i).Type
		for elem.Kind() == reflect.Ptr || elem.Kind() == reflect.Slice || elem.Kind() == reflect.Array {
			elem = elem.Elem()
		}
		if elem.Kind() != reflect.Struct && elem.Kind() != reflect.Interface {
			continue
		}
		if found >= 0 {
			return reflect.Value{}, false
		}
		found = i
	}
	if found < 0 {
		return reflect.Value{}, false
	}
	return unwrapValue(v.Field(found)), true
}

// unwrapValue peels interface and pointer wrappers so reflection sees the
// concrete value. Needed because paginated results are collected as
// []interface{} and SDK models are often pointers.
func unwrapValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return v
		}
		v = v.Elem()
	}
	return v
}

func printSliceAsTable(v reflect.Value) error {
	if v.Len() == 0 {
		return nil
	}

	// Get the first element to determine structure. Paginated commands collect
	// into []interface{}, so the interface has to be unwrapped before the element
	// looks like a struct — otherwise every list falls through to the bullet
	// fallback and prints raw Go struct dumps instead of a table.
	first := unwrapValue(v.Index(0))

	if first.Kind() != reflect.Struct {
		// Not a slice of structs, print as list
		for i := 0; i < v.Len(); i++ {
			item := unwrapValue(v.Index(i))
			fmt.Printf("%s %v\n", IconBullet, item.Interface())
		}
		return nil
	}

	// Build table from slice of structs
	headers, rows := structSliceToTable(v)

	// Browsing shows every column and pans, so it skips narrowing entirely.
	if browseOutput && BrowseFunc != nil {
		shown, err := BrowseFunc(headers, rows)
		if err != nil {
			return err
		}
		if shown {
			return nil
		}
		// Not a terminal: fall through to static output.
	}

	if !wideOutput {
		idx := SelectColumns(headers, rows, activePreferred, TerminalWidth())
		headers, rows = projectColumns(headers, rows, idx)
	}

	PrintTable(headers, rows)
	return nil
}

// projectColumns reduces headers and rows to the given column indices.
func projectColumns(headers []string, rows [][]string, idx []int) ([]string, [][]string) {
	if len(idx) == 0 || len(idx) == len(headers) {
		return headers, rows
	}
	h := make([]string, 0, len(idx))
	for _, i := range idx {
		h = append(h, headers[i])
	}
	out := make([][]string, 0, len(rows))
	for _, row := range rows {
		r := make([]string, 0, len(idx))
		for _, i := range idx {
			if i < len(row) {
				r = append(r, row[i])
			} else {
				r = append(r, "")
			}
		}
		out = append(out, r)
	}
	return h, out
}

func printStructAsTable(v reflect.Value) error {
	t := v.Type()

	fmt.Println(Styles.Title.Render(t.Name()))

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		name := jsonFieldName(field)

		// Handle pointer values
		displayValue := value
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				fmt.Printf("  %s: %s\n",
					Styles.Bold.Render(name),
					Styles.Muted.Render("-"))
				continue
			}
			displayValue = value.Elem()
		}

		fmt.Printf("  %s: %v\n",
			Styles.Bold.Render(name),
			displayValue.Interface())
	}

	return nil
}

func printMapAsTable(v reflect.Value) error {
	keys := v.MapKeys()

	for _, key := range keys {
		value := v.MapIndex(key)
		fmt.Printf("  %s: %v\n",
			Styles.Bold.Render(fmt.Sprintf("%v", key.Interface())),
			value.Interface())
	}

	return nil
}

func structSliceToTable(v reflect.Value) ([]string, [][]string) {
	if v.Len() == 0 {
		return nil, nil
	}

	// Get struct type from first element
	first := unwrapValue(v.Index(0))
	t := first.Type()

	// Build headers from struct fields
	var headers []string
	var fieldIndices []int

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		name := jsonFieldName(field)

		headers = append(headers, name)
		fieldIndices = append(fieldIndices, i)
	}

	// Build rows
	var rows [][]string
	for i := 0; i < v.Len(); i++ {
		item := unwrapValue(v.Index(i))
		// A heterogeneous slice would panic on Field(); skip anything that is not
		// the struct type the headers were derived from.
		if item.Kind() != reflect.Struct || item.Type() != t {
			continue
		}

		var row []string
		for _, idx := range fieldIndices {
			value := item.Field(idx)
			if value.Kind() == reflect.Ptr {
				if value.IsNil() {
					row = append(row, "-")
					continue
				}
				value = value.Elem()
			}
			row = append(row, fmt.Sprintf("%v", value.Interface()))
		}
		rows = append(rows, row)
	}

	return headers, rows
}
