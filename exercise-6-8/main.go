package main

import "fmt"

func main() {
	fmt.Println("hello")
	f := foo()

	f()
}

func foo() func() {
	return func() {
		fmt.Println("i'm a returned function")
	}
}
