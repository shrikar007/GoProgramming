package main
import "fmt"

func main(){

num :=  ([]int {1,3,5,7})
num1 := ([]int {2,4,6,8})

num = append(num,num1...)

fmt.Println(num)


num = append  (num[:2],num[3:]...)//deleting an element

fmt.Println(num)

}
