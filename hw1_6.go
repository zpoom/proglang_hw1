package main

import "fmt"

func incrementByValue(x int){
	x++
}
func incrementByPointer(x *int){
	*x++
}
func main(){
	var x = 0
	fmt.Println("main : ",x)
	incrementByValue(x)
	fmt.Println("after increment by value : ",x)
	incrementByPointer(&x)
	fmt.Println("after increment by pointer : ",x)
}