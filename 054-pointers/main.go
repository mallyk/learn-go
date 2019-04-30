package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //hives address
	fmt.Printf("%T\n", a)

	b := &a
	fmt.Printf("%T\n", &a)
	fmt.Println(*b) //gives you teh value stored at an address
}
