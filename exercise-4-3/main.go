package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for k, v := range x {
		fmt.Println(k, v)
	}

	fmt.Printf("%T\n", x)

	y := x[:5]
	z := x[5:]
	a := x[2:7]
	b := x[1:6]

	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
}
