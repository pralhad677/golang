package main 

import (
	"fmt"
)

type myStruct struct {
	name string
}

func (m myStruct) getName() string {
	fmt.Println(m)
	return m.name
}
func (m *myStruct)  getName1() string {
	fmt.Println(m)
	
	return m.name
}

func  main()  {
	m:= myStruct{name:"jacob"}
	fmt.Println(m.getName())
	fmt.Println(m.getName1())
	
}