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
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			counter++
			fmt.Println("counter value in process:", counter)
			fmt.Println("goroutines:\t", runtime.NumGoroutine())
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("final counter value:", counter)
	fmt.Println("fin")
}
