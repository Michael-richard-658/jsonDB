package main

import (
	"fmt"
	"log"

	compiler "github.com/Michael-richard-658/Simple-database/Compiler"
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
	//for {
	query := `create table bikes(
	model text,
	hp int,
	nm int,
	ABS boolean
	);`
	Tokens := Compiler.Lexer(query)
	AST, err := Compiler.Parser(Tokens)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("AST: %+v\n", AST)
	/*if err != nil {
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
