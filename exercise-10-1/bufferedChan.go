package main

import "fmt"

func main() {
	fmt.Println("hello")

	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
