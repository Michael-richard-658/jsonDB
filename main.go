package main

import (
	"github.com/Michael-richard-658/Simple-database/operationdb"
)

func main() {
	DBCRUD := operationdb.UserCRUD{}
	//DBCRUD.CreateTable("BIKES", " NAME CC HP TORQUE ")
	/*DBCRUD.InsertRecord("BIKES", `NAME: Xpulse 200,
	CC: 199.5,
	HP: 20.5,
	TORQUE: 18.1`,
	)*/
	DBCRUD.QueryRecord("SELECT * FROM BIKES ")
}
