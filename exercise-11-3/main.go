package main

import "fmt"

type customErr struct {
}

func (c customErr) Error() string {
	return "error occured"
}

func main() {
	cError := customErr{}
	foo(cError)
}

func foo(e error) {
	fmt.Println(e)
}
