package main

import (
	"fmt"
)
type arr []string

func main(){
	a :=[]string{"a","b"}
	for i,item:=range a{
		fmt.Println(i,item);
	}
	fmt.Println("hi")
}