package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}
 
func (p Person) addAge(num int) int{
	return p.age+num
}


func main(){
	p1:=Person{name:"jacob",age:23}
		
	

	fmt.Println(p1.name)
	fmt.Println(p1.addAge(2))

}