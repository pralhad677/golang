package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	g:=100
	var wg sync.WaitGroup
	wg.Add(g * 2)
	n:=0
	
	var m sync.Mutex
	//mutex comes from mutual exclusion
	//code betwn m.lock and m.unlock are only executed by one goroutine
	for i:=0;i<g;i++{
			go func(){
				time.Sleep(time.Second /10)
				m.Lock()
				n++
				m.Unlock()
				wg.Done()
			}()

			go func(){
				time.Sleep(time.Second /10)
				m.Lock()
				n--
				m.Unlock()
				wg.Done()
			}()
	}
	wg.Wait()
	fmt.Println("final value of n",n)
	fmt.Println("jacob")
}

//to detect race condn use -race on command like go run -race  main.go