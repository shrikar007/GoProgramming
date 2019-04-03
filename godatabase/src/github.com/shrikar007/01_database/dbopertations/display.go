package dbopertations

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/shrikar007/01_database/structure"
	"log"
	"net/http"
)

func Display(w http.ResponseWriter, r *http.Request) {
	Db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")
	var (
		id int
		fname, lname string
	)
	rows, err := Db.Query("select * from student")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &fname, &lname)
		if err != nil {
			log.Fatal(err)
		}
		message := structure.Msg{
			id,fname,lname}

		data, err := json.Marshal(message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, string(data))
	}
}

