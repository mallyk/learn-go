package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter\t", atomic.LoadInt64(&counter))
			//time.Sleep(time.Second)
			runtime.Gosched()
			wg.Done()
			fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
		}()
	}

	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("count:", counter)
}
