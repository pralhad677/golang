package main

import (
	"fmt"
)
type arr []string

func main(){
	var a arr=[]string{"a","b"}
	for i,item:=range a{
		fmt.Println(i,item);
	}
	fmt.Println("hi")
}