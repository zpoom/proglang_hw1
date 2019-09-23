package main

import "fmt"

func incrementByValue(x int){
	x++
}
func incrementByReference(y []int){
	y[0]++
	y[1]++
	y[2]++
}
func main(){
	var x = 0
	var y = []int{0,1,2}
	y[0] = 0
	y[1] = 1
	y[2] = 2
	fmt.Println("main : ",x)
	incrementByValue(x)
	fmt.Println("after call incrementByValue : ",x)
	incrementByReference(y)
	fmt.Println("after call incrementByReference : ",y)
}