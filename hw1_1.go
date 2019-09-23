package main

import "fmt"
var x = 0
func foo() {
	var x = 2
	fmt.Println("foo : ",x)
}
func main(){
	var x = 1
	fmt.Println("main : ",x)
	foo()
}