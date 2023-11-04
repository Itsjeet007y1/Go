package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	counter := 0

	const gs = 10000

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Go Routines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go Routines", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)
}
