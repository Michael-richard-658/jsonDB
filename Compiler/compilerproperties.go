package compiler

type CompilerPropertiesInterface interface {
	Lexer(query string) []Token
}
type CompilerProperties struct{}
