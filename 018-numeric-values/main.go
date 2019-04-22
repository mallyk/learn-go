package main

import "fmt"

var x int
var y float64
var z uint8 //byte is alias for uint8

//rune is alias for int32

func main() {
	x = 2
	y = 42.33534
	z = 127
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
