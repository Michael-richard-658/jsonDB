package main

import (
	"fmt"
	"log"

	compiler "github.com/Michael-richard-658/Simple-database/Compiler"
	engine "github.com/Michael-richard-658/Simple-database/Engine"
	"github.com/Michael-richard-658/Simple-database/utils"
	"github.com/chzyer/readline"
)

func main() {
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

	Compiler := compiler.CompilerProperties{}
	Engine := engine.EngineProperties{}
	//for {

	queries := []string{`create table bikes(
	model text,
	hp int,
	nm int,
	ABS boolean
	);`,
		"select * from bikes;"}
	Tokens := Compiler.Lexer(queries[0])
	AST, err := Compiler.Parser(Tokens)
	if err != nil {
		fmt.Println(err)
		return
	}
	queryPlan := Engine.QueryPlanner(AST)
	Engine.QueryExecutor(queryPlan)
	/*if err != nil {}
		break
	}

	query = strings.TrimSpace(query)
	if query == "" {
		continue
	}
	if strings.ToLower(query) == "exit" {
		fmt.Println("Bye!")
		break
	}*/

	//}
}
