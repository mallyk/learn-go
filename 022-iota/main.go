package main

import "fmt"

const (
	a = iota //automatically increment by `1`
	b = iota
	c = iota
)

const (
	d = iota //resets with another const declaration
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
