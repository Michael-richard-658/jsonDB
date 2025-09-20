package operationdb

type DBoperations struct{}
type DataBaseCRUD interface {
	QueryParser(query string) []string
	CreateTable(tableName string, attributes string)
	QueryRecord(query []string)
	InsertRecord(tableName string, query string)
	UpdateData()
	DeleteData()
	DropTable(tableName string)
	DescTable(query string)
}
