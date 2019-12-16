package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	mylock RWLock
	myerr  = errors.New("my error")
)

type RWLock struct {
	lock sync.RWMutex
}

// WithRLock runs the given function with the read lock grabbed.
func (rwm *RWLock) WithRLock(f func() error) error {
	(&rwm.lock).RLock()
	defer (&rwm.lock).RUnlock()
	return f()
}

// WithWLock runs the given function with the write lock grabbed.
func (rwm *RWLock) WithWLock(f func() error) error {
	(&rwm.lock).Lock()
	defer (&rwm.lock).Unlock()
	return f()
}

func retchan() (<-chan int, error) {
	notifyCh := make(chan int, 5*10)

	var err error
	for i := 1; i <= 5; i++ {
		err = mylock.WithWLock(func() error {
			// if i == 4 {
			// 	return myerr
			// }

			fmt.Printf("notifyCh <- %d \n", i)
			notifyCh <- i
			return nil
		})
		if err != nil {
			close(notifyCh)
			return nil, err
		}

		// up := 1
		// notifyCh <- up
	}

	return notifyCh, nil
}

func main() {
	notiCh, err := retchan()
	if err != nil {
		fmt.Println(err)
	}

	timer := time.NewTimer(time.Duration(1*time.Second) + time.Duration(rand.Int63n(int64(1*time.Second))))
	for {
		select {
		case <-timer.C:
			fmt.Printf("Auto disconnect with client, timer.C: %+v \n", timer.C)
			return
		case updates, ok := <-notiCh:
			if !ok {
				// Channel has been closed.
				fmt.Println("Channel has been closed.")
				return
			}

			fmt.Printf("notfiCh %+v\n", updates)
		}
	}
}
