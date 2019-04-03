package dbopertations

import (
	"database/sql"

	"log"
)
func Delete(id int){
	Db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")

	_,err=Db.Query("delete  from student where id=?",id)
	if err!=nil{
		log.Fatal(err)
	}
}