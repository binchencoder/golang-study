package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			fmt.Println("Automatic ticker.")
		}
	}()

	fmt.Println("End")

	wg.Wait()
}
