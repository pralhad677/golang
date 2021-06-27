// Converting Go objects into JSON is called marshalling in Golang

package main

import (
	"fmt"
	"encoding/json"	
)
type Employee struct {
	Name  string
	Age	int	
	Address string
}

func main(){
	e:=Employee{Name:"jacob",Age:23,Address:"ktm"}

	fmt.Println(e)
	data:=fmt.Sprintf("object is,%v",e)
	fmt.Println(data)
	empData,err:= json.Marshal(e) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(empData)	//buffer data
	// fmt.Scanf("type of empData is %T ",empData)
	fmt.Println(string(empData))	
}