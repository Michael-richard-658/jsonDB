package main

import (
	"github.com/Michael-richard-658/Simple-database/operationdb"
)

func main() {
	DBCRUD := operationdb.UserCRUD{}
	//DBCRUD.CreateTable("DOGS", " NAME BREED  AGE COLOR ")
	/*DBCRUD.InsertRecord("DOGS", `NAME: COOKESH,
	 BREED: ST-BERNARD, AGE: 5, COLOR: WHITE-BROWN`,
	)*/
	DBCRUD.QueryRecord("SELECT * FROM BIKES ")
}
