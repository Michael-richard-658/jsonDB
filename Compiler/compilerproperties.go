package compiler

type CompilerPropertiesInterface interface {
	Lexer(query string) []Token
	Parser(tokens []Token) (interface{}, error)
	QueryPlanner(ast AST) interface{}
	QueryExecutor(plan interface{})
}
type CompilerProperties struct{}
