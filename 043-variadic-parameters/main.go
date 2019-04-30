package main

import "fmt"

func main() {
	sum := foo(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("the total is: ", sum)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding,", v, "to the total which is now,", sum)
	}

	return sum
}
