package main

import (
	"fmt"
)

type Gift struct {
	name string
	price float64
}

func getGift(g Gift){
	g.name = "shirt"
	g.price = 200.00
}
func getGiftWithPointer(g *Gift){
	(*g).name = "shirt"
	(*g).price = 200.00
}

func main(){
	g:=Gift{name:"pant",price:1300}
	getGift(g)
	fmt.Println("without pointer")
	fmt.Printf("name of gift %v \n",g.name)
	fmt.Printf("price of gift %v \n",g.price)
	fmt.Println("with pointer")
	getGiftWithPointer(&g)
	fmt.Printf("name of gift %v \n",g.name)
	fmt.Printf("price of gift %v \n",g.price)
}

//note
// only int float struct string bool  cab be modified by pointers