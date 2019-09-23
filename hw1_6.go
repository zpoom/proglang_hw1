package main

import "fmt"

func incrementByReference(x *int){
	*x++
}

func incrementByValue(x int){

}

func main(){
	var x = 0
	fmt.Println("main : ",x)
	incrementByValue(x)
	fmt.Println("after call incrementByValue : ",x)
	incrementByReference(&x)
	fmt.Println("after call incrementByReference : ",x)
}