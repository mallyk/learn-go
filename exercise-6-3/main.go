package main

import "fmt"

func main() {

	foo()
}

func foo() {
	defer bar()
	fmt.Println("fair enough")
}

func bar() {
	fmt.Println("i ran after")
}
