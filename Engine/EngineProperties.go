package engine

import (
	compiler "github.com/Michael-richard-658/Simple-database/Compiler"
)

type EngineProperties struct{}
type EnginePropertiesInterface interface {
	QueryPlanner(ast compiler.AST) interface{}
	QueryExecutor(plan interface{})
}
