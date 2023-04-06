package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Goroutine 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 4; i++ {
			mu.Lock()
			fmt.Println("[coba1 coba2 coba3]", i)
			mu.Unlock()
		}
	}()

	// Goroutine 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 4; i++ {
			mu.Lock()
			fmt.Println("[bisa1 bisa2 bisa3]", i)
			mu.Unlock()
		}
	}()

	wg.Wait()
}
