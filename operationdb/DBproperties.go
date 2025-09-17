package operationdb

type DataBaseCRUD interface {
	CreateTable(tableName string, attributes string)
	InsertRecord(tableName string, query string)
	UpdateData()
	DeleteData()
	QueryRecord(query string)
}

type UserCRUD struct{}
