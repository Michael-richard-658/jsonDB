package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Michael-richard-658/Simple-database/operationdb"
	"github.com/Michael-richard-658/Simple-database/utils"
	"github.com/chzyer/readline"
)

//!!!Move query parser func to operationdb package
/*
!! pass query to query parser and to seprate and and
create table function instead of current
*/
func main() {
	DBCRUD := operationdb.UserCRUD{}
	utilFunctions := utils.Utils{}
	utilFunctions.Clrscr()

	queryLine, err := readline.NewEx(&readline.Config{
		Prompt:          "jsonDB> ",
		HistoryFile:     "/tmp/db_history.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer queryLine.Close()

	for {
		query, err := queryLine.Readline()
		if err != nil {
			break
		}

		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		if strings.ToLower(query) == "exit" {
			fmt.Println("Bye!")
			break
		}

		queryParts := utilFunctions.QueryParser(query)
		actionType := strings.ToUpper(queryParts[0])
		switch actionType {
		case "CREATE":
			tableName, fields := operationdb.MakeCreateTableQueryArray(query)
			DBCRUD.CreateTable(tableName, strings.Join(fields, " "))
		case "SELECT":
			DBCRUD.QueryRecord(query)
		case "INSERT":
			fmt.Println("Performing INSERT operation")
		case "UPDATE":
			fmt.Println("Performing UPDATE operation")
		case "DELETE":
			fmt.Println("Performing DELETE operation")
		case "DESC":
			DBCRUD.DescTable(queryParts)
		default:
			fmt.Println("Invalid Query, please check syntax!")
		}
	}
}
