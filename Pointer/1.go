package main

import (
	"fmt"
)

type Person struct {
	age int
}

func (p Person) b(x *int){
	fmt.Println("memory Address of x",x)
	
	*x = 5
	// x=324
	fmt.Println("memory Address of x",x)
}

func main(){
	a :=3
	fmt.Println(a)
	p :=Person{age:23}
	p.b(&a)
	fmt.Println(a)
	
	var ptr1 *float64 // * in front of type is type description
	//dereferencing 
	*p //pointer le point gareko place ko data

}

// note 
//&value->pointer
//*pointer->value