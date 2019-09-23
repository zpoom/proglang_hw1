package main

import "fmt"


func normalForLoop(){
	// normal for loop
	for i:=1 ; i < 5 ;i++ {
		fmt.Println("Hello world #",i)
	}
}

func forLoopAsWhile(){
	// use for loop as while loop
	i:= 1
	for i < 5 {
		fmt.Println("Hello world #",i)
		i++
	}
}
func forLoopAsDoWhile(){
	i:=1
	for {
		fmt.Println("Hello world #",i)
		i++
		// break condition
		if i >= 5{
			break
		}
	}
}
func ifCondition(){
	var x = 2019
	if x % 2 == 0 {
		fmt.Println("x is even.")
	}else {
		fmt.Println("x is odd.")
	}
}
func switchCase(){
	var x = 2019
	switch x{
		case 2562: 
			fmt.Println("x is B.E.")
		case 2019: 
			fmt.Println("x is A.D.")
		default: 
			fmt.Println("This year is not x")
	}
}
func main(){
	switchCase()
	// ifCondition()
	// normalForLoop()
	// forLoopAsWhile()
	// forLoopAsDoWhile() // if remove break condition it will be an infinite loop

}