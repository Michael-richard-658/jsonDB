package compiler

import (
	"strings"
)

type TokenType string

const (
	SELECT TokenType = "SELECT"
	FROM   TokenType = "FROM"
	CREATE TokenType = "CREATE"
	WHERE  TokenType = "WHERE"

	IDENTIFIER TokenType = "IDENTIFIER"
	NUMBER     TokenType = "NUMBER"
	EQUAL      TokenType = "EQUAL"
	COMMA      TokenType = "COMMA"
	SEMICOLON  TokenType = "SEMICOLON"
)

type Token struct {
	Type  TokenType
	Value string
}

var keywords = map[string]TokenType{
	"SELECT": SELECT,
	"FROM":   FROM,
	"CREATE": CREATE,
	"WHERE":  WHERE,
}

func (cp *CompilerProperties) Lexer(query string) []Token {
	var tokens []Token
	currentToken := ""

	for i := 0; i < len(query); i++ {
		ch := query[i]

		if ch == ' ' || ch == ',' || ch == ';' || ch == '=' {
			if currentToken != "" {
				tokens = append(tokens, classifyWord(currentToken))
				currentToken = ""
			}
			if ch == ',' {
				tokens = append(tokens, Token{Type: COMMA, Value: ","})
			}
			if ch == ';' {
				tokens = append(tokens, Token{Type: SEMICOLON, Value: ";"})
			}
			if ch == '=' {
				tokens = append(tokens, Token{Type: EQUAL, Value: "="})
			}
			continue
		}

		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') || ch == '*' {
			currentToken += string(ch)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, classifyWord(currentToken))
	}

	return tokens
}

func classifyWord(word string) Token {
	if isNumber(word) {
		return Token{Type: NUMBER, Value: word}
	}
	if tok, ok := keywords[strings.ToUpper(word)]; ok {
		return Token{Type: tok, Value: word}
	}
	return Token{Type: IDENTIFIER, Value: word}
}

func isNumber(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}
