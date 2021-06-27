//pointer to pointer
package main

import (
	"fmt"
)


func main(){
		a:=3
		p1:=&a
		//pointer to pointer
		p2:=&p1

		
		fmt.Println("for p1")
		fmt.Printf("type of p1 is %T and value is %v\n",p1,p1 )
		fmt.Printf("memory Address of p1 %p\n",&p1 )
		fmt.Printf("*p1 is %v\n",*p1 )

		fmt.Println("for p2")
		fmt.Printf("type of p2 is %T and value is %v\n",p2,p2 )
		fmt.Printf("memory Address of p2 %p\n",&p2 )
		fmt.Printf("*p2 %v\n",*p2 )
		fmt.Printf("**p2 %v\n",**p2 )

}