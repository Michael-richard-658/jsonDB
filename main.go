package main

import (
	"github.com/Michael-richard-658/Simple-database/operationdb"
)

func main() {
	DBCRUD := operationdb.UserCRUD{}
	//DBCRUD.CreateTable("BIKES", " NAME CC HP TORQUE ")
	DBCRUD.InsertRecord("DOGS", `NAME: Royal Enfield Classic 500,
	CC: 499,
	HP:27.,
	TORQUE: 41.3`,
	)
}
