package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("hello")
	counter := 0
	gs := 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("counter value in process:", counter)
			fmt.Println("goroutines:\t", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("final counter value:", counter)
	fmt.Println("fin")
}
