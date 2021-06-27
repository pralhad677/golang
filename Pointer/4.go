package main


import (
	"fmt"
)

func k(a *int) *float64{
	*a =5
	b:=6.5
	return &b
}

func main(){
	y:=new(int)
	a:=3
	y=&a
	fmt.Printf("value of a before change %v \n",a)
	k(y)
	
	fmt.Printf("value of a after change %v",a)
}