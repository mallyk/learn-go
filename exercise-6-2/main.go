package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(s...))
	fmt.Println(bar(s))
}

func foo(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}

	return sum
}

func bar(vi []int) int {
	var sum int
	for _, v := range vi {
		sum += v
	}

	return sum
}
