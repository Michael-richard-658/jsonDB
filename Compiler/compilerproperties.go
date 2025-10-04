package compiler

type CompilerPropertiesInterface interface {
	Lexer(query string) []Token
	Parser(tokens []Token) (interface{}, error)
}
type CompilerProperties struct{}
