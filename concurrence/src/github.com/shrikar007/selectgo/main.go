package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
//time.Sleep(5 * time.Second)
ch <- "from server1"
}
func server2(ch chan string) {
//time.Sleep(3* time.Second)
ch <- "from server2"

}
func main() {
//output1 := make(chan string)
//output2 := make(chan string)

ch:=make(chan string,2)
go server2(ch)
go server1(ch)
time.Sleep(1*time.Second)
select {
case s1 := <-ch:
fmt.Println(s1)
case s2 := <-ch:
fmt.Println(s2)

}

fmt.Println("The End")
}
