package main 

import (
	"fmt"
)

func changeSlice(arr []int){
	
	for i,item:=range arr{
		fmt.Println(i,item);
		arr[i]++
		fmt.Println(arr[i]);
		// fmt.Println(arr[item]);
	}
}

func changeMap(m map[string]int){
	m["a"] = 2
	m["b"] = 3
	m["c"] = 4

}

func main(){

	array:=[]int{1,2,3,4}
	
	fmt.Println(array) //[1 2 3 4 ]
	changeSlice(array)
	fmt.Println(array) //[2 3 4 5]

	// map:=map[string]int{"a":1,"b":2,"c":3}
	// fmt.Println(map)
	
	map1 :=make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	
	fmt.Println("after calling change map")
	changeMap(map1)
	fmt.Println(map1)
}

//note slice and map already contains a pointer to its key so we do not need pointer to cahnge its conatent  