package main

import (
	"fmt"
)

func main(){
	map1 :=make(map[int]string)
	map1[1] = "one"
	map1[2] = "two"
	fmt.Println(map1[1])
	fmt.Println(map1[2])
}