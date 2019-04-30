package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
n	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
