package engine

import (
	compiler "github.com/Michael-richard-658/Simple-database/Compiler"
)

func (ep *EngineProperties) QueryPlanner(ast compiler.AST) interface{} {
	switch ast.StatementType {
	case "CREATE":
		return createTablePlanner(ast)

	case "SELECT":
	}
	return nil
}

// CREATE PLANNING
type CreateTablePlan struct {
	Operation string
	TableName string
	Columns   []compiler.ColumnDef
}

func createTablePlanner(ast compiler.AST) *CreateTablePlan {
	return &CreateTablePlan{
		Operation: "CREATE_TABLE",
		TableName: ast.ASTNode.(*compiler.CreateTableStmt).TableName,
		Columns:   ast.ASTNode.(*compiler.CreateTableStmt).Columns,
	}
}
