package output

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

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

var currentFormat = FormatAuto

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
func Print(data interface{}) error {
	format := currentFormat

	// Auto-detect format
	if format == FormatAuto {
		if IsTerminal() {
			format = FormatTable
		} else {
			format = FormatJSON
		}
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
		return printStructAsTable(v)

	case reflect.Map:
		return printMapAsTable(v)

	default:
		// Fall back to JSON for unknown types
		return printJSON(data)
	}
}

func printSliceAsTable(v reflect.Value) error {
	if v.Len() == 0 {
		return nil
	}

	// Get the first element to determine structure
	first := v.Index(0)
	if first.Kind() == reflect.Ptr {
		first = first.Elem()
	}

	if first.Kind() != reflect.Struct {
		// Not a slice of structs, print as list
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			if item.Kind() == reflect.Ptr {
				item = item.Elem()
			}
			fmt.Printf("%s %v\n", IconBullet, item.Interface())
		}
		return nil
	}

	// Build table from slice of structs
	headers, rows := structSliceToTable(v)
	PrintTable(headers, rows)
	return nil
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

		// Get JSON tag name if available
		name := field.Name
		if tag := field.Tag.Get("json"); tag != "" && tag != "-" {
			// Parse json tag (handle omitempty, etc.)
			for j := 0; j < len(tag); j++ {
				if tag[j] == ',' {
					name = tag[:j]
					break
				}
			}
			if name == "" {
				name = tag
			}
		}

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
	first := v.Index(0)
	if first.Kind() == reflect.Ptr {
		first = first.Elem()
	}
	t := first.Type()

	// Build headers from struct fields
	var headers []string
	var fieldIndices []int

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		// Get JSON tag name
		name := field.Name
		if tag := field.Tag.Get("json"); tag != "" && tag != "-" {
			for j := 0; j < len(tag); j++ {
				if tag[j] == ',' {
					name = tag[:j]
					break
				}
			}
			if name == "" {
				name = tag
			}
		}

		headers = append(headers, name)
		fieldIndices = append(fieldIndices, i)
	}

	// Build rows
	var rows [][]string
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
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
