package operationdb

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// check not done yet
func (u *DBoperations) DropTable(queryParts []string) {
	if len(queryParts) < 3 {
		fmt.Println("invalid drop command")
		return
	}
	tableName := strings.ToUpper(queryParts[2])
	filePath := filepath.Join("./tables", tableName+".json")
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("failed to delete table  %s \n", tableName)
		return
	}
	println("table dropped successfully")
}
