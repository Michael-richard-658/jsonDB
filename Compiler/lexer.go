package compiler

import (
	"strings"
)

type TokenType string

const (
	SELECT      TokenType = "SELECT"
	CREATE      TokenType = "CREATE"
	DESC        TokenType = "DESC"
	FROM        TokenType = "FROM"
	WHERE       TokenType = "WHERE"
	EOF         TokenType = "EOF"
	IDENTIFIER  TokenType = "IDENTIFIER"
	NUMBER      TokenType = "NUMBER"
	EQUAL       TokenType = "EQUAL"
	PARENTHESIS TokenType = "PARENTHESIS"
	COMMA       TokenType = "COMMA"
	SEMICOLON   TokenType = "SEMICOLON"
	INT         TokenType = "INT"
	TEXT        TokenType = "TEXT"
	BOOLEAN     TokenType = "BOOLEAN"
)

type Token struct {
	Type  TokenType
	Value string
}

var keywords = map[string]TokenType{
	"SELECT": SELECT,
	"DESC":   DESC,
	"FROM":   FROM,
	"CREATE": CREATE,
	"WHERE":  WHERE,
	"EOF":    EOF,
}
var dataTypes = map[string]TokenType{
	"INT":     INT,
	"TEXT":    TEXT,
	"BOOLEAN": BOOLEAN,
}

func (cp *CompilerProperties) Lexer(query string) []Token {
	var tokens []Token
	currentToken := ""

	for i := 0; i < len(query); i++ {
		ch := query[i]

		if ch == ' ' || ch == ',' || ch == ';' || ch == '=' || ch == '(' || ch == ')' {
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
			if ch == '(' {
				tokens = append(tokens, Token{Type: PARENTHESIS, Value: "("})
			}
			if ch == ')' {
				tokens = append(tokens, Token{Type: PARENTHESIS, Value: ")"})
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
	upperWord := strings.ToUpper(word)

	if isNumber(word) {
		return Token{Type: NUMBER, Value: word}
	}

	if tok, ok := keywords[upperWord]; ok {
		return Token{Type: tok, Value: word}
	}

	if tok, ok := dataTypes[upperWord]; ok {
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
