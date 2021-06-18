package main

import (
	"fmt"
	"time"
)
func a(ch chan int){
	a:=5 
	ch <- a
	fmt.Println("ends")
}

func main(){
	channel:=make(chan int)
	defer close(channel)
	go a(channel)
	time.Sleep(time.Second *2)
	num:=<- channel
	
	fmt.Println(num)
	
	
	fmt.Println("jacob")
	fmt.Printf("%T",channel)
}