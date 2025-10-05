package compiler

import (
	"fmt"
	"strings"
)

// === TOKEN HANDLING HELPERS ===

var tokenIndex int

func current(tokens []Token) Token {
	if tokenIndex >= len(tokens) {
		return Token{Type: EOF, Value: ""}
	}
	return tokens[tokenIndex]
}

func next() {
	tokenIndex++
}

func expect(expected TokenType, tokens []Token) (Token, error) {
	if current(tokens).Type == expected {
		tok := current(tokens)
		next()
		return tok, nil
	}
	return Token{}, fmt.Errorf("expected %v, got %v", expected, current(tokens).Type)
}

// === AST STRUCTURES ===

type AST struct {
	ASTNode       interface{}
	StatementType string
}

// === MAIN PARSER ===

func (cp *CompilerProperties) Parser(tokens []Token) (AST, error) {
	tokenIndex = 0

	switch tokens[0].Type {
	case SELECT:
		stmt, err := parseSelectStmt(tokens)
		return AST{ASTNode: stmt, StatementType: "SELECT"}, err
	case CREATE:
		stmt, err := parseCreateStmt(tokens)
		return AST{ASTNode: stmt, StatementType: "CREATE"}, err
	case DESC:
		stmt, err := parseDescStmt(tokens)
		return AST{ASTNode: stmt, StatementType: "DESC"}, err
	default:
		return AST{}, fmt.Errorf("unexpected start of statement: %v", tokens[0].Value)
	}
}

// === CREATE TABLE PARSER ===

type ColumnDef struct {
	Name   string
	Type   string
	Length int
}

type CreateTableStmt struct {
	TableName string
	Columns   []ColumnDef
}

func parseCreateStmt(tokens []Token) (*CreateTableStmt, error) {
	if len(tokens) < 8 {
		return nil, fmt.Errorf("invalid CREATE TABLE statement")
	}

	if strings.ToUpper(tokens[1].Value) != "TABLE" {
		return nil, fmt.Errorf("expected 'TABLE' after CREATE, got %s", tokens[1].Value)
	}

	if tokens[3].Value != "(" {
		return nil, fmt.Errorf("expected '(' after table name")
	}

	stmt := &CreateTableStmt{TableName: tokens[2].Value}
	var columns []ColumnDef
	i := 4
	expectColumn := true

	for i < len(tokens) {
		token := tokens[i]

		if token.Value == ")" || token.Type == SEMICOLON {
			break
		}

		if expectColumn {
			if token.Type != IDENTIFIER {
				return nil, fmt.Errorf("expected column name, got '%s'", token.Value)
			}
			colName := token.Value
			i++

			if i >= len(tokens) {
				return nil, fmt.Errorf("expected data type for column '%s'", colName)
			}

			colType := strings.ToUpper(tokens[i].Value)
			colLength := 0

			// Handle VARCHAR(n)
			if colType == "VARCHAR" {
				if i+3 >= len(tokens) ||
					tokens[i+1].Value != "(" ||
					tokens[i+2].Type != NUMBER ||
					tokens[i+3].Value != ")" {
					return nil, fmt.Errorf("invalid VARCHAR declaration for column '%s'. Correct format: VARCHAR(n)", colName)
				}
				colLength = atoi(tokens[i+2].Value)
				i += 3 // skip '(', number, ')'
			} else if colType != "INT" && colType != "BOOLEAN" {
				return nil, fmt.Errorf("invalid data type '%s' for column '%s'", colType, colName)
			}

			columns = append(columns, ColumnDef{
				Name:   colName,
				Type:   colType,
				Length: colLength,
			})

			expectColumn = false
			i++
		} else {
			if token.Type != COMMA {
				return nil, fmt.Errorf("expected ',' between columns, got '%s'", token.Value)
			}
			expectColumn = true
			i++
		}
	}

	if expectColumn && len(columns) > 0 {
		return nil, fmt.Errorf("dangling comma at end of column list")
	}
	if i >= len(tokens) || tokens[i].Value != ")" {
		return nil, fmt.Errorf("missing closing parenthesis")
	}
	if tokens[len(tokens)-1].Type != SEMICOLON {
		return nil, fmt.Errorf("missing ';' at end of statement")
	}

	stmt.Columns = columns
	return stmt, nil
}

func atoi(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

// === SELECT PARSER ===

type Condition struct {
	Left     string
	Operator string
	Right    any
}

type SelectStmt struct {
	Columns []string
	Table   string
	Where   *Condition
}

func parseSelectStmt(tokens []Token) (*SelectStmt, error) {
	tokenIndex = 0
	_, err := expect(SELECT, tokens)
	if err != nil {
		return nil, err
	}

	cols, err := parseColumnList(tokens)
	if err != nil {
		return nil, err
	}

	_, err = expect(FROM, tokens)
	if err != nil {
		return nil, err
	}

	tableTok, err := expect(IDENTIFIER, tokens)
	if err != nil {
		return nil, err
	}

	var where *Condition
	if current(tokens).Type == WHERE {
		next()
		where, err = parseCondition(tokens)
		if err != nil {
			return nil, err
		}
	}

	_, err = expect(SEMICOLON, tokens)
	if err != nil {
		return nil, err
	}

	return &SelectStmt{
		Columns: cols,
		Table:   tableTok.Value,
		Where:   where,
	}, nil
}

func parseColumnList(tokens []Token) ([]string, error) {
	var cols []string

	first, err := expect(IDENTIFIER, tokens)
	if err != nil {
		return nil, err
	}
	cols = append(cols, first.Value)

	for current(tokens).Type == COMMA {
		next()
		tok, err := expect(IDENTIFIER, tokens)
		if err != nil {
			return nil, err
		}
		cols = append(cols, tok.Value)
	}
	return cols, nil
}

func parseCondition(tokens []Token) (*Condition, error) {
	left, err := expect(IDENTIFIER, tokens)
	if err != nil {
		return nil, err
	}

	op, err := expect(EQUAL, tokens)
	if err != nil {
		return nil, err
	}

	right, err := expect(IDENTIFIER, tokens)
	if err != nil {
		return nil, err
	}

	return &Condition{
		Left:     left.Value,
		Operator: op.Value,
		Right:    right.Value,
	}, nil
}

// === DESC PARSER ===

type DescStmt struct {
	TableName string
}

func parseDescStmt(tokens []Token) (*DescStmt, error) {
	if len(tokens) < 2 {
		return nil, fmt.Errorf("incomplete DESC statement")
	}
	return &DescStmt{TableName: tokens[1].Value}, nil
}
