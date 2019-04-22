package main

import "fmt"

//var declaration can be used in and out of function
//scope of Z is package level

//DECLARE the variable z
//ASSIGN the VALUE 43
// declare & assign = initialization
var z = 44

//declare there is a variable with identifier zz
//and that the variable is of TYPE int
//ASSIGNS the ZERO VALUE of TYPE int to zz
var zz int

func main() {
	//DECLARE a variable and assign VALUE
	x := 42
	fmt.Println(x)
	var y = 43
	fmt.Println(y)
	fmt.Println(z)
	foo()
	fmt.Println(zz)
}

func foo() {
	fmt.Println("again:", z)
}
