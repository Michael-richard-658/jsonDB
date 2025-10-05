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
	VARCHAR     TokenType = "VARCHAR"
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
	"VARCHAR": VARCHAR,
	"BOOLEAN": BOOLEAN,
}

// Lexer converts a SQL string into tokens
func (cp *CompilerProperties) Lexer(query string) []Token {
	var tokens []Token
	currentToken := ""

	isNumber := func(s string) bool {
		for _, ch := range s {
			if ch < '0' || ch > '9' {
				return false
			}
		}
		return true
	}

	flush := func() {
		if currentToken != "" {
			tok := classifyWord(currentToken, isNumber)
			tokens = append(tokens, tok)
			currentToken = ""
		}
	}

	for i := 0; i < len(query); i++ {
		ch := query[i]
		switch ch {
		case ' ', '\t', '\n', '\r':
			flush()
		case ',', ';', '=', '(', ')':
			flush()
			switch ch {
			case ',':
				tokens = append(tokens, Token{Type: COMMA, Value: ","})
			case ';':
				tokens = append(tokens, Token{Type: SEMICOLON, Value: ";"})
			case '=':
				tokens = append(tokens, Token{Type: EQUAL, Value: "="})
			case '(':
				tokens = append(tokens, Token{Type: PARENTHESIS, Value: "("})
			case ')':
				tokens = append(tokens, Token{Type: PARENTHESIS, Value: ")"})
			}
		default:
			currentToken += string(ch)
		}
	}

	flush() // flush any remaining token
	return tokens
}

// classifyWord determines the token type for a word
func classifyWord(word string, isNumber func(string) bool) Token {
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
