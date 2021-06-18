package main

import (
	"fmt"
	"sync"
	"time"
)
func a(c1 chan int,wg *sync.WaitGroup){
	fmt.Println("2:Execution of a goroutine started")
	c1 <-1
	wg.Done()
	fmt.Println("4:Execution of a goroutine ended")
}

func main(){
	fmt.Println("1:Execution of main goroutine started")
	var wg sync.WaitGroup
	c1:=make(chan int,1) //2nd parameter chai capacitor ho
	defer close(c1)
	wg.Add(1)	

	go a(c1,&wg)
	time.Sleep(time.Second*2)
	fmt.Println("3:data from channel",<-c1)
	
	
	wg.Wait()
	fmt.Println("5:Execution of main goroutine Stopped")
}

func random(){
	fmt.Println("jacob")
}

//output
// 1:Execution of main goroutine started
// 2:Execution of a goroutine started
// 3:data from channel 
// 4:Execution of a goroutine ended
// 5:Execution of main goroutine Stopped


//malai tha nai avyena k aru tw k vanam k
//my name is jacob ryan
//my name is jacob ryan
//my name is jacob ryan









