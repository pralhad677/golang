package main

import (
	"fmt"
	"errors"
)

func Hello(name string)(string,error){
	if name=="" {
		return "",
		errors.New("empty Name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}

func main(){
	fmt.Println("jacob")
	name:=""
	fmt.Println(Hello(name))
}

//main difference is in go we can do multiple return but in js we can return only one

//typescript 	equivalent
// interface Func<T> {
//     (x:string):T|Error
// }
// let a:Func<string> =(x:string) =>{
//     if(x.length ===0) {
//       throw new Error('String is empty')
//     }
//     return x

// }

// let b = a('')
// console.log(b)