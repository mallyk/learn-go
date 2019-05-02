package main

import "fmt"

func main() {
	fmt.Println("hello")

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
