package operationdb

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/Michael-richard-658/Simple-database/utils"
)

func (u *UserCRUD) InsertRecord(tableName string, query string) {
	fullPath := filepath.Join("./tables", tableName+".json")

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Fatalf("Table %s does not exist!", tableName)
		return
	}

	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open table %s: %v", tableName, err)
	}
	defer file.Close()

	data, err := os.ReadFile(fullPath)
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
	fmt.Println("Sorted keys from query:", sortedKeys)

	content := string(data)
	stringToJson, err := utils.StringToJSON(content)
	if err != nil {
		log.Fatal("String to JSON error! ")
	}

	for _, key := range sortedKeys {
		if _, ok := stringToJson[key]; ok {
			fmt.Println("Table matched for key:", key)
		} else {
			log.Fatalf("Wrong table name or unknown field: %s", key)
		}
	}

	/*
		_, err = file.Write([]byte("\n" + query + "\n"))
		if err != nil {
			fmt.Println("Something went wrong while writing to the file, please try again")
			return
		}
	*/
}
