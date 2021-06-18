package main

import (
	"fmt"
	"errors"
)

func Hello(name string)(string,error){
	if name=="" {
		return "",
		errors.New("empty Name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}

func main(){
	fmt.Println("jacob")
	name:=""
	fmt.Println(Hello(name))


	
}