package main

import (
	"fmt"
	"runtime"
	
	"sync"
)

func f1(wg *sync.WaitGroup){
	fmt.Println("execution of go routine started")
	slice:=[]int{1,2,3,4,5}
	for i,item:=range slice{
		fmt.Println(i,item)
	}
	wg.Done()
	fmt.Println("execution of go routine end")
	// (*wg).Done()
}

func f2(){
	fmt.Println("second function")
}

func main(){

	var wg sync.WaitGroup
	//wg.Add(1) means for wait for 1 go routine
	wg.Add(1)
	fmt.Println("main execution of program started")
	fmt.Println("number of cpu ",runtime.NumCPU())
	fmt.Println("number of goroutines ",runtime.NumGoroutine())
	fmt.Println("OS ",runtime.GOOS)
	fmt.Println("Arch ",runtime.GOARCH)
	fmt.Println("GOMAXPROCS ",runtime.GOMAXPROCS(0))
	go f1(&wg)
	fmt.Println("number of goroutines after second go routine started ",runtime.NumGoroutine())

	f2()
	// time.Sleep(time.Second*1)
	wg.Wait()
	fmt.Println("end of execution of main")
}
