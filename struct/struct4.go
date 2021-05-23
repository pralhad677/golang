//example with pointer receiver

package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

func (p *Person) fn(name1 string) {
	p.name = name1
}

func main(){
	p1:=&Person{name:"jacob",age:23}

	p1.fn("ryan")
	fmt.Println(p1.name);

}