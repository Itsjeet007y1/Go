package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Bigin CPU:", runtime.NumCPU())
	fmt.Println("Bigin GRTN:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from main 1!!")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from main 2!!")
		wg.Done()
	}()

	fmt.Println("Mid CPU:", runtime.NumCPU())
	fmt.Println("Mid GRTN:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End CPU:", runtime.NumCPU())
	fmt.Println("End GRTN:", runtime.NumGoroutine())
	fmt.Println("about to exit")
}
