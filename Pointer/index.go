package main

import (
	"fmt"
)

func b(x *int){
	fmt.Println(x)
	c:=6
	*x = c
	fmt.Println(x)
	// fmt.Println(&x)

}


func main(){
	a :=3;
	fmt.Println(&a)
	fmt.Println(a)
	b(&a)
	fmt.Println(a)
	// b(a) error occurs
}