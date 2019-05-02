package main

import "fmt"

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for val := 0; val < 10; val++ {
				c <- val
			}
		}()
	}

	for v := 0; v < 100; v++ {
		fmt.Println(v, <-c)
	}

	fmt.Println("about to exit")
}
