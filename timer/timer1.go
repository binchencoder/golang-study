package main

import (
	"fmt"
	"time"
)

func main() {
	go timer1()
	go timer2()
	fmt.Println("will end")
	time.Sleep(5 * time.Second)
}

func testTimer1() {
	go func() {
		fmt.Println("test timer1")
	}()
}

func testTimer2() {
	go func() {
		fmt.Println("test timer2")
	}()
}

func timer1() {
	timer1 := time.NewTimer(2 * time.Second)
	select {
	case <-timer1.C:
		testTimer1()
	}
}

func timer2() {
	timer2 := time.NewTicker(2 * time.Second)
	select {
	case <-timer2.C:
		testTimer2()
	}
}
