package operationdb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Michael-richard-658/Simple-database/utils"
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
	result = append(result, parts[0])
	result = append(result, attributes)
	for _, v := range parts[i:] {
		result = append(result, v)
	}

	return result
}

func (u *UserCRUD) QueryRecord(query string) {
	queryParts := makeQueryArray(query)

	fromIndex := -1
	for i, part := range queryParts {
		partsToString, ok := part.(string)
		if ok && strings.ToUpper(partsToString) == "FROM" && i+1 < len(queryParts) {
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

	var records []map[string]any
	if err := json.Unmarshal(fileData, &records); err != nil {
		log.Fatalf("Failed to unmarshal table %s: %v", tableName, err)
	}

	if len(records) == 0 {
		log.Println("Table is empty")
		return
	}
	utilsDB := utils.Utils{}
	//utilsDB.MapToJSON(records[0], "obj")
	utilsDB.MapToSQLTable(records, "*")
}
