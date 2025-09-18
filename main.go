package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Michael-richard-658/Simple-database/operationdb"
	"github.com/Michael-richard-658/Simple-database/utils"
)

func main() {
	DBCRUD := operationdb.UserCRUD{}
	utilsFunctions := utils.Utils{}
	utilsFunctions.Clrscr()
	run := true
	for run {
		reader := bufio.NewReader(os.Stdin)
		print("mysql> ")
		query, _ := reader.ReadString('\n')
		queryParts := utilsFunctions.QueryParts(query)
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
		default:
			fmt.Println("Invalid Query please check synatx!")
		}
	}
}
