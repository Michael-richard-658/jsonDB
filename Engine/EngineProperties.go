package engine

import (
	compiler "github.com/Michael-richard-658/Simple-database/Compiler"
)

type EnginePropertiesInterface interface {
	QueryPlanner(ast compiler.AST) interface{}
	QueryExecutor(plan interface{})
}
type EngineProperties struct{}
