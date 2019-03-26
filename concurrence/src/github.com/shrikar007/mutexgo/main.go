package main

import (
	"fmt"
	"sync"
	"time"
)

var x=0
func increament1(m *sync.Mutex){
     m.Lock()
     x++

     m.Unlock()
	fmt.Println(x)
}

//func increament2(){

//	x++

//}

func main(){

	var m sync.Mutex

	go increament1(&m)
	go increament1(&m)
	time.Sleep(10000)
	//fmt.Println(x)

}
