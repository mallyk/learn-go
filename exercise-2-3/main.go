package main

import "fmt"

const (
	x string = "blargh"
	y        = 32
)

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
