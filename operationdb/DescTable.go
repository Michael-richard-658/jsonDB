package operationdb

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Michael-richard-658/Simple-database/utils"
)

func (u *UserCRUD) DescTable(queryParts []string) {
	tableNameWithSemicolon := queryParts[1]
	tableName := strings.TrimSuffix(tableNameWithSemicolon, ";")
	fullPath := filepath.Join("./tables", strings.ToUpper(tableName)+".json")
	fileData, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("Failed to read table %s: %v", tableName, err)
	}
	var records []map[string]any
	if err := json.Unmarshal(fileData, &records); err != nil {
		log.Fatalf("Invalid JSON in table %s: %v", tableName, err)
	}
	dbUtils := utils.Utils{}
	dbUtils.MapToSQLTable(records[0], "2")
}
