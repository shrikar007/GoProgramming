
package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

const port = 7070

type Msg struct {
	Data string `json:"data"`
}

func main() {
	http.HandleFunc("/helloworld", myHandler)
	log.Printf("Starting server on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	message := Msg{Data: "Hello, World!"}

	//data, err := json.Marshal(message)
	encoder := json.NewEncoder(w)
	encoder.Encode(&message)

	//if err != nil {
		//log.Fatal(err)
	//}

	//fmt.Fprint(w, data)
}
