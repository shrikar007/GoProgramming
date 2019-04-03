package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shrikar007/01_database/dbopertations"
	"net/http"
)
func main(){
	var id,ch int
	var fname,lname string

	fmt.Println("1.insert:\n2.display:\n3:delete:\nEnter your choice=")
	fmt.Println("Enter choice=")
	fmt.Scan(&ch)
	switch ch{
	case 1:		fmt.Println("Enter Student data=")
				fmt.Scan(&id,&fname,&lname)
				dbopertations.Insert(id,fname,lname)


	case 2:
				http.HandleFunc("/display", dbopertations.Display)
				http.ListenAndServe(":6011",nil)


	case 3:		fmt.Println("Enter id to delete:")
				fmt.Scan(&id)
	        	dbopertations.Delete(id)

	default:	fmt.Println("invallid choice")

	}
}


