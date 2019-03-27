package main

import (
	"log"
	"net/http"
	"testing"
	"io/ioutil"
	"encoding/json"
	"fmt"
)
var port1=7171
func main() {
	http.HandleFunc("/helloworld", myHandler)
	log.Printf("Starting server on port %v", port1)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port1), nil))
}

type Msg1 struct {
	Data string `json:"data"`
}

func BenchmarkHandlerMarshal(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	message := Msg1{"Hello, World!"}
	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(message)
		fmt.Fprint(writer, string(data))
	}
}

func BenchmarkHandlerEncoder(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	message := Msg1{"Hello, World!"}
	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(message)
	}
}

func BenchmarkHandlerEncoderReference(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	message := Msg1{"Hello, World!"}
	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&message)
	}
}
