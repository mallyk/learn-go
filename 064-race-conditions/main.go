package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
		}()
	}

	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("count:", counter)
}
