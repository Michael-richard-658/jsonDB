package operationdb

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func (u *UserCRUD) InsertRecord(tableName string, query string) {
	fullPath := filepath.Join("./tables", tableName+".json")

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Fatalf("Table %s does not exist!", tableName)
	}

	fileData, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("Failed to read table %s: %v", tableName, err)
	}

	re := regexp.MustCompile(`\b([A-Za-z_]+)\s*:`)
	matches := re.FindAllStringSubmatch(query, -1)
	sortedKeys := []string{}
	for _, match := range matches {
		sortedKeys = append(sortedKeys, match[1])
	}
	sort.Strings(sortedKeys)

	var records []map[string]interface{}
	if err := json.Unmarshal(fileData, &records); err != nil {
		log.Fatalf("Invalid JSON in table %s: %v", tableName, err)
	}
	if len(records) == 0 {
		log.Fatalf("Table %s has no schema definition", tableName)
	}

	schema := records[0]
	keyCount := 0
	for _, key := range sortedKeys {
		if _, ok := schema[key]; ok {
			keyCount++
		} else {
			log.Fatalf("Unknown field: %s", key)
		}
	}

	if keyCount == len(schema) {
		recordMap := make(map[string]interface{})
		parts := strings.Split(query, ",")
		for _, part := range parts {
			pair := strings.SplitN(strings.TrimSpace(part), ":", 2)
			if len(pair) != 2 {
				continue
			}
			key := strings.TrimSpace(pair[0])
			val := strings.TrimSpace(pair[1])
			if i, err := strconv.Atoi(val); err == nil {
				recordMap[key] = i
			} else if f, err := strconv.ParseFloat(val, 64); err == nil {
				recordMap[key] = f
			} else {
				recordMap[key] = val
			}
		}

		records = append(records, recordMap)
		jsonData, err := json.MarshalIndent(records, "", "    ")
		if err != nil {
			log.Fatalf("Failed to marshal records: %v", err)
		}

		if err := os.WriteFile(fullPath, jsonData, 0644); err != nil {
			log.Fatalf("Failed to write table %s: %v", tableName, err)
		}
	} else if keyCount > len(schema) {
		log.Fatal("More fields in query than in table definition.")
	}
}
