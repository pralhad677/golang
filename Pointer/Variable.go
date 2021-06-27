package main

import (
	"fmt"

)

//pointer pass by vale vs pass by reference

func shopping(pant string,price float64,available bool){
	//pointer napathdua chai each args lai memory ko xuattai palce ma copy garerko ho and tesma modify gareko ho original ma haina just like by value in js just like local variable environment
	pant= "Levis Denim"
	price = 1800.00
	available = true
}
func shoppingWithPointer(pant *string,price *float64,available *bool){
	//pointer napathdua chai each args lai memory ko xuattai palce ma copy garerko ho and tesma modify gareko ho original ma haina just like by value in js just like local variable environment
	*pant= "Levis Denim"
	*price = 1800.00
	*available = true
}

func main(){
		pant,price,available := "Fila Denim",2400.00,false
		fmt.Println("before calling shopping function",)
		fmt.Printf("pant %v \n",pant)
		fmt.Printf("price %v \n",price)
		fmt.Printf("available %v \n",available)
		shopping(pant,price,available)
		fmt.Println("after calling shopping function",)
		fmt.Printf("pant %v \n",pant)
		fmt.Printf("price %v \n",price)
		fmt.Printf("available %v \n",available)
	
		fmt.Println("now calling shopping with pointer")
		shoppingWithPointer(&pant,&price,&available)
		fmt.Printf("pant %v \n",pant)
		fmt.Printf("price %v \n",price)
		fmt.Printf("available %v \n",available)		
}