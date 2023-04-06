package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// Goroutine 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 4; i++ {
			fmt.Println("[bisa1 bisa2 bisa3]", i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()

	// Goroutine 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 4; i++ {
			fmt.Println("[coba1 coba2 coba3]", i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()

	wg.Wait()
}
