package utils

import (
	"encoding/json"
	"path/filepath"
)

func CreateFilePath(tableName string) string {
	dirPath := "./tables"
	fileExtension := ".json"

	// filePath will change based on the tableName passed into the function
	filePath := filepath.Join(dirPath, tableName+fileExtension)
	return filePath
}

func StringToJSON(str string) (map[string]any, error) {
	var result map[string]any
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
