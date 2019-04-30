package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()

	fmt.Printf("%T\n", x)
	fmt.Println(x())     //x being ran
	fmt.Println(bar()()) //same thing without being assigned to x
}

func foo() string {
	return "hello world"
}

func bar() func() int {
	return func() int {
		return 451
	}
}
