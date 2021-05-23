package main

import (
	"fmt"
)
type Address struct {
    Name    string
    city    string
    Pincode int
}
 

func main(){
	// a: =Address{"ryan","ktm",123}
	// fmt.Println(a.Name) //error dinxa access grnalai tala ko jhai lehna prxa
	
	//object le . operator ko through bata property access grna lagda class lai dart ma jhai bane Paramete type ma lekhna prxa
    a1 := Address{Name:"jacob", city:"Kathmandu",Pincode: 234234}
	
	fmt.Println(a1.Name)
	fmt.Println(a1.city)
	fmt.Println("hi")
}