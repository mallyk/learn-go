package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("hello")
	var counter int64
	gs := 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("counter value in process:", atomic.LoadInt64(&counter))
			fmt.Println("goroutines:\t", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("final counter value:", counter)
	fmt.Println("fin")
}
