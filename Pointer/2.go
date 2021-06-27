package main

import (
	"fmt"
)

func main(){
	p :=new(int) //decalring pointer of type int
	x:=3
	p=&x
	fmt.Printf("type of p is %T and value is %v\n",p,p )
	fmt.Printf("memory Address of x %p\n",&x )

	*p=5 //equivalent to x =5
	fmt.Println(*p) //5
	fmt.Println(x) //5
	fmt.Println("*p == x",*p==x) //true

}