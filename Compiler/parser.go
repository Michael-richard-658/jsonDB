package compiler

import "fmt"

// Where STRUCT
type Condition struct {
	Left     string
	Operator string
	Right    any
}

// HELPERS FOR PARSING
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

func (cp *CompilerProperties) Parser(tokens []Token) (interface{}, error) {

	stmt, err := parseStatement(tokens)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}
func parseStatement(tokens []Token) (interface{}, error) {
	switch current(tokens).Type {
	case SELECT:
		return parseSelectStmt(tokens)
	case CREATE:
		return nil, nil
	case DESC:
		return parseDescStmt(tokens)
	default:
		return nil, fmt.Errorf("unexpected token %v", current(tokens))
	}
}

// SELECT PARSING
type SelectStmt struct {
	Columns []string
	Table   string
	Where   *Condition
}

func parseSelectStmt(tokens []Token) (*SelectStmt, error) {
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

// DESC PARSING
type DescStmt struct {
	Table Token
}

func parseDescStmt(tokens []Token) (*DescStmt, error) {
	return &DescStmt{
		Table: tokens[1],
	}, nil
}
