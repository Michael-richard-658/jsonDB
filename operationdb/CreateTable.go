package operationdb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (u *DBoperations) CreateTable(tableName string, attributes string) {
	dirPath := "./tables"

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("failed to read directory: %s", err)
	}

	for _, entry := range entries {
		if entry.Name() == strings.ToUpper(tableName)+".json" {
			fmt.Printf("Table with name %v already exists.\n", tableName)
			return
		}
	}

	attributeArr := strings.Fields(attributes)
	getLastAttribute := attributeArr[len(attributeArr)-1]
	if getLastAttribute[len(getLastAttribute)-1] == ';' {
		attributeArr[len(attributeArr)-1] = strings.TrimSuffix(getLastAttribute, ";")
	}
	record := make(map[string]string)
	for _, attr := range attributeArr {
		record[attr] = ""
	}

	jsonData, err := json.MarshalIndent(record, "", "    ")
	if err != nil {
		log.Fatalf("failed to marshal map: %v", err)
	}

	filePath := filepath.Join(dirPath, strings.ToUpper(tableName)+".json")

	createFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("failed to create table: %s", err)
	}
	defer createFile.Close()
	_, err = createFile.WriteString("[" + string(jsonData) + "]")
	if err != nil {
		fmt.Println("Something went wrong while writing to the file, please try again")
		return
	}

	fmt.Printf("Table %v created successfully!\n", tableName)
}
