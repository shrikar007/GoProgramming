package dbopertations

import (
	"database/sql"
	"fmt"

)

func Insert (id int,fname,lname string) {
	Db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")

	stmt,err:=Db.Prepare("insert into student values (?,?,?)")
	if err!=nil{
		fmt.Println(err)
	}
	stmt.Exec(id,fname,lname)
	stmt.Close()

}