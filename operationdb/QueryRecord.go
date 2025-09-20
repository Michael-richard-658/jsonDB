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

func (u *DBoperations) QueryRecord(queryParts []string) {

	fromIndex := -1
	for i, part := range queryParts {

		if strings.ToUpper(part) == "FROM" && i+1 < len(queryParts) {
			fromIndex = i + 1
			break
		}
	}

	if fromIndex == -1 {
		log.Fatal("Invalid query: missing FROM")
	}

	tableNameWithSemicolon := queryParts[fromIndex]
	tableName := strings.ToUpper(strings.TrimSuffix(tableNameWithSemicolon, ";"))
	fullPath := filepath.Join("./tables", fmt.Sprintf("%s.json", tableName))

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		fmt.Printf("Table %s does not exist! \n", tableName)
		return
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
	if len(records) > 1 {
		utilsDB.MapToSQLTable(records, "*")
	} else {
		println("Empty set!")
	}
}
