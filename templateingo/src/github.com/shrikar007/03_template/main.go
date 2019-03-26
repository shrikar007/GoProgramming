package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func  main() {
	fname := os.Args[1]
	lname:=os.Args[2]

	template := `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Welcome to Go Templating</title>
		</head>
		<body>
			<h1>` + fname + `</h1>
            <h1>` + lname + `</h1>
		</body>
	</html>
	`
	f, err := os.Create("name.html")
	fmt.Printf("%T",f)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	io.Copy(f, strings.NewReader(template))
}

