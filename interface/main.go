package main

import (
	"fmt"
)
type P interface  {
	getName() string
	getAge() int
}

type Person struct{
	name string
	age int
}
func (p Person) getName() string {
	return p.name
}
func (p Person) getAge() int {
	return p.age
}

func a(i P) string{
return	i.getName()
}
func main(){
	p:=Person{name:"jacob",age:23}
	fmt.Println(p.name)
	fmt.Println("hey")
	fmt.Println(a(p))

}