package main

import "database/sql"

func main(){
	var db *sql.DB = connDB()
	queryData(&db,"000001")
}
