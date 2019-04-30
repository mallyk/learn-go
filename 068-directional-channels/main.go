package main

import "fmt"

func main() {
	fmt.Println("hello")

	c := make(chan int, 1)
	cr := make(<-chan int) //receive
	cs := make(chan<- int) //send

	fmt.Println("-------------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	//general to specific is ok
	cr = c
	cs = c

	//specifc to general does not work
	// c = cr
	// c = cs

	fmt.Println("-------------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

}
