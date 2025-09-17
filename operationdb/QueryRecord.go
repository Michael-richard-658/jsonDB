package operationdb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func makeQueryArray(query string) []any {
	parts := strings.Fields(query)

	if parts[1] == "*" {
		result := []any{}
		for _, p := range parts {
			result = append(result, p)
		}
		return result
	}

	attributes := []string{}
	i := 1
	for i < len(parts) && strings.ToUpper(parts[i]) != "FROM" {
		attributes = append(attributes, parts[i])
		i++
	}

	result := []any{}
	result = append(result, parts[0])   // "SELECT"
	result = append(result, attributes) // nested fields
	for _, v := range parts[i:] {
		result = append(result, v)
	}

	return result
}

func (u *UserCRUD) QueryRecord(query string) {
	queryParts := makeQueryArray(query)

	// find table name after FROM
	fromIndex := -1
	for i, v := range queryParts {
		s, ok := v.(string)
		if ok && strings.ToUpper(s) == "FROM" && i+1 < len(queryParts) {
			fromIndex = i + 1
			break
		}
	}

	if fromIndex == -1 {
		log.Fatal("Invalid query: missing FROM")
	}

	tableName, ok := queryParts[fromIndex].(string)
	if !ok {
		log.Fatal("Invalid table name")
	}

	fullPath := filepath.Join("./tables", fmt.Sprintf("%s.json", tableName))

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Fatalf("Table %s does not exist!", tableName)
	}

	fileData, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("Failed to read table %s: %v", tableName, err)
	}

	// Unmarshal into a slice of maps
	var records []map[string]interface{}
	if err := json.Unmarshal(fileData, &records); err != nil {
		log.Fatalf("Failed to unmarshal table %s: %v", tableName, err)
	}

	if len(records) == 0 {
		log.Println("Table is empty")
		return
	}

	// Print schema keys (first record)
	fmt.Println("Schema keys:", records[0])

	// Optional: print all records
	for i, record := range records {
		fmt.Printf("Record %d: %+v\n", i+1, record)
	}
}
