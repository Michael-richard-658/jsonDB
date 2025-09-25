package main

import (
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

	//for {
	Compiler := compiler.CompilerProperties{}
	query := "select cc,hp from users  ;"
	Tokens := Compiler.Lexer(query)
	Compiler.Parser(Tokens)
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
