package main

import (
	"fmt"
	"sync"
	"time"
)

func a(c chan int,wg *sync.WaitGroup)  {
	fmt.Println("2:execution a of goroutine started")
	c <- 3
	wg.Done()
	fmt.Println("4:execution a of goroutine ended")

}
//unbuffered means euta(A) goroutine le arko(B) goroutine lai channel ko through data pass grda  B go routine le channel ko 
//data napayesamma execution rokinxa A ko


func main(){
	fmt.Println("1:Execution of main started")
	var wg sync.WaitGroup
	wg.Add(1)
	c1:=make(chan int) // unbuffered channel
	defer close(c1)
	go a(c1,&wg)
	time.Sleep(time.Second*2)
	fmt.Println("3:channel value",<-c1)
	wg.Wait()
	// fmt.Println("channel value",<-c1) deadlock occurs  using channel after wg.Wait()
	fmt.Println("5:Execution of main ended")

}
//output
// 1:Execution of main  started
// 2:Execution of a goroutine started
// 3:channel value
// 4:Execution of a goroutine ended
// 5:Execution of main  Stopped


////asndkjasbdkas