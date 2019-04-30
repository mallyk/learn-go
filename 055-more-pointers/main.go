package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x before foo", x, &x)
	foo(x)
	fmt.Println("x after foo", x, &x)

	fmt.Println("x before bar", x, &x)
	bar(&x)
	fmt.Println("x after bar", x, &x)
}

func foo(y int) {
	fmt.Println("y before foo", y)
	y = 43
	fmt.Println("y after foo", y)
}

func bar(y *int) {
	fmt.Println("y before bar", *y)
	*y = 43
	fmt.Println("y after bar", *y)
}
