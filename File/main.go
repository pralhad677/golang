package main

import (
	// "bufio"
    "fmt"
    // "io"
    "io/ioutil"
    // "os"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main(){
	dat, err := ioutil.ReadFile("./a.txt")
    check(err)

	// f, err := os.Open("/tmp/dat")
    // check(err)
	
    fmt.Print(string(dat))
    // fmt.Print(string(f))

}