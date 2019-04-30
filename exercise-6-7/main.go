package main

import "fmt"

func main() {
	f := foo

	f()
}

func foo() {
	fmt.Println("i've been assigned")
}
