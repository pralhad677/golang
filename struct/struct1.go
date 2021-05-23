package main

import (
	"fmt"
)



type Creature struct {
    Name     string
    Greeting string
}

func (c Creature) Greet() {
    fmt.Printf("%s says %s", c.Name, c.Greeting)
}

interface Person struct {
	name string
	age int
}

func (p Person) addAge(num int){
	x :=p.age+num
    fmt.Printf(x)
}

func main() {
    sammy := Creature{
        Name:     "Sammy",
        Greeting: "Hello!",
    }
    sammy.Greet() //instance like in js
	Creature.Greet(sammy) //static like in js

	person:= Person{
		name:"jacob",
		age:23
	}

	Person.addAge(person,1);
	person.addAge(1);
}