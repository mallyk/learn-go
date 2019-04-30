package main

import "fmt"

func main() {

	foo()
	bar("Kyle")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// func ( r receiver) identifier(parameters) (return(s)){...}

func foo() {
	fmt.Println("hello from foo")
}

//everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("hello from woo,", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, "says hello")
	b := false
	return a, b
}
