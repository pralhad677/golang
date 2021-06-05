package main

import (
	"fmt"
	"math/rand"	
	"time"
)


func Random(n int) int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(n)
	return value
}
func getRandom(c chan int){
	data :=Random(10)
	c <- data
}

func main(){
	intChaneel :=make(chan int)
	defer close(intChaneel)
	go getRandom(intChaneel)
	dataFromChannel :=<-intChaneel

	fmt.Println(dataFromChannel)
	fmt.Println("my name is jacob")
}