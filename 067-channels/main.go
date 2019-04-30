package main

import "fmt"

func main() {
	c := make(chan int) //without buffer

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func successfulBuffer() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

func unsuccessfulBuffer() {
	c := make(chan int, 1)

	c <- 42
	c <- 43

	fmt.Println(<-c)
}
