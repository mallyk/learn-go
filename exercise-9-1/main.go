package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("main go routine")

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("go routine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("go routine 2")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("fin")
}
