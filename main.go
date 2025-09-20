package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Michael-richard-658/Simple-database/operationdb"
	"github.com/Michael-richard-658/Simple-database/utils"
	"github.com/chzyer/readline"
)

//!!!Insert query
//!! select query update on only fields required
/*
! remove fatal from killing program return errors instead
! make a separate file for queryparser which comes under DBcompiler package
*/
func main() {
	DBOperation := operationdb.DBoperations{}
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

		queryParts, attributes := DBOperation.QueryParser(query)

		actionType := strings.ToUpper(queryParts[0])
		switch actionType {
		case "CREATE":
			tableName := queryParts[2]
			DBOperation.CreateTable(tableName, strings.Join(attributes, " "))
		case "SELECT":
			DBOperation.QueryRecord(queryParts)
		case "INSERT":
			fmt.Println("Performing INSERT operation")
		case "UPDATE":
			fmt.Println("Performing UPDATE operation")
		case "DELETE":
			fmt.Println("Performing DELETE operation")
		case "DESC":
			DBOperation.DescTable(queryParts)
		case "DROP":
			DBOperation.DropTable(queryParts)
		default:
			fmt.Println("Invalid Query, please check syntax from main!")
		}
	}
}
