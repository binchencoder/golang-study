package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock sync.Mutex
)

func ilock() {
	lock.Lock()
	time.Sleep(10 * time.Millisecond)
}

func timerTest() {
	// ilock()

	lock.Lock()
	defer lock.Unlock()

	eps := 1

	notifyCh := make(chan int)

	timer := time.NewTimer(1 * time.Millisecond)
	for {
		select {
		case <-timer.C:
			fmt.Println("Timer expired")
			// return
		case notifyCh <- eps:
			fmt.Println(notifyCh)
			if !timer.Stop() {
				<-timer.C
			}
		}
	}

	fmt.Println(notifyCh)
}

func main() {
	// go timerTest()

	timerTest()
}
