package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1(){
	slice:=[]int{1,2,3,4,5}
	for i,item:=range slice{
		fmt.Println(i,item)
	}
	fmt.Println("execution of go routine started")
}

func f2(){
	fmt.Println("second function")
}

func main(){
	fmt.Println("main execution of program started")
	fmt.Println("number of cpu ",runtime.NumCPU())
	fmt.Println("number of goroutines ",runtime.NumGoroutine())
	fmt.Println("OS ",runtime.GOOS)
	fmt.Println("Arch ",runtime.GOARCH)
	fmt.Println("GOMAXPROCS ",runtime.GOMAXPROCS(0))
	go f1()
	fmt.Println("number of goroutines after second go routine started ",runtime.NumGoroutine())

	f2()
	time.Sleep(time.Second*1)
	fmt.Println("end of execution of main")
}
