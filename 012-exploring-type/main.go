package main

import "fmt"

var y = 42

//a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
//DECLARED the VARIABLE with the IDENTIFIER z is of TYPE string and ASSIGN the value "Shaken, not stirred"
var z string = "Shaken, not stirred"
var a string = `James said`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
