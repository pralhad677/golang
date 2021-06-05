package main

import (
	"fmt"
)
type arr []string

func main(){
	var arr []string
	arr =append(arr,"ryan")
	arr =append(arr,"jacob")
	for i,item:=range arr{
		fmt.Println(i,item);
	}
	fmt.Println("hi")
}