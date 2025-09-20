package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type Utils struct{}
type UtilsInterface interface {
	MapToSQLTable(data interface{}, mode string)
	//	MapToJSON(m map[string]any, mode string) string
	Clrscr()
}

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}

func printBorder(keys []string, widths map[string]int) {
	b := "+"
	for _, k := range keys {
		b += strings.Repeat("-", widths[k]+2) + "+"
	}
	fmt.Println(b)
}

func mapsToSQLTable(rows []map[string]interface{}, skipFirst bool) {
	if len(rows) == 0 {
		fmt.Println("No rows to display")
		return
	}

	keys := make([]string, 0, len(rows[0]))
	for k := range rows[0] {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	widths := make(map[string]int, len(keys))
	for _, k := range keys {
		widths[k] = len(k)
	}
	for _, r := range rows {
		for _, k := range keys {
			val := fmt.Sprintf("%v", r[k])
			if len(val) > widths[k] {
				widths[k] = len(val)
			}
		}
	}

	printBorder(keys, widths)

	h := "|"
	for _, k := range keys {
		h += " " + padRight(k, widths[k]) + " |"
	}
	fmt.Println(h)
	printBorder(keys, widths)

	start := 0
	if skipFirst {
		start = 1
	}

	for _, r := range rows[start:] {
		line := "|"
		for _, k := range keys {
			line += " " + padRight(fmt.Sprintf("%v", r[k]), widths[k]) + " |"
		}
		fmt.Println(line)
		printBorder(keys, widths)
	}
}

func (u Utils) MapToSQLTable(data any, mode string) {
	switch mode {
	case "1":
		if m, ok := data.(map[string]interface{}); ok {
			mapsToSQLTable([]map[string]interface{}{m}, false)
		} else {
			fmt.Println("Invalid type for mode 1: expected map[string]interface{}")
		}
	case "*":
		if arr, ok := data.([]map[string]interface{}); ok {
			mapsToSQLTable(arr, true)
		} else {
			fmt.Println("Invalid type for mode *: expected []map[string]interface{}")
		}
	case "2":
		if m, ok := data.(map[string]interface{}); ok {
			keys := make([]string, 0, len(m))
			for k := range m {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			widths := make(map[string]int, len(keys))
			for _, k := range keys {
				widths[k] = len(k)
			}

			printBorder(keys, widths)
			h := "|"
			for _, k := range keys {
				h += " " + padRight(k, widths[k]) + " |"
			}
			fmt.Println(h)
			printBorder(keys, widths)
		}
	default:
		fmt.Println("Unknown mode. Use \"1\" for single map, \"*\" for slice of maps, or \"2\" for schema only.")
	}
}
func (u Utils) MapToJSON(m map[string]any, mode string) string {
	jsonData, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Fatalf("failed to marshal map: %v", err)
	}

	if mode == "print" {
		fmt.Println(string(jsonData))
		return ""
	}

	if mode == "obj" {
		return string(jsonData)
	}

	log.Fatalf("invalid mode: %s (use 'print' or 'obj')", mode)
	return ""
}

func (u *Utils) Clrscr() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

}
