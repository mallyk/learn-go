package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("anonymous func ran")
	}()

	func(x int) {
		fmt.Println("the meaning of life:", x)
	}(42)
}

func foo() {
	fmt.Println("foo ran")
}
