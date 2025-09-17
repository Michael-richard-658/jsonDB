package operationdb

type DataBaseCRUD interface {
	CreateTable(tableName string, attributes string)
	InsertRecord(tableName string, query string)
	UpdateData()
	DeleteData()
	QueryRecord()
}

type UserCRUD struct{}
