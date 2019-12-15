package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendMessage(sc chan string) {
	var i int

	for {
		i = rand.Intn(10)
		for ; i >= 0; i-- {
			sc <- fmt.Sprintf("Other number %d", rand.Intn(100))
		}
		i = 1000 + rand.Intn(32000)
		time.Sleep(time.Duration(i) * time.Millisecond)
	}
}

func sendNum(c chan int) {
	var i int
	for {
		i = rand.Intn(16)
		for ; i >= 0; i-- {
			time.Sleep(20 * time.Millisecond)
			c <- rand.Intn(65534)
		}
		i = 1000 + rand.Intn(24000)
		time.Sleep(time.Duration(i) * time.Millisecond)
	}
}

func main() {
	msgchan := make(chan string, 32)
	numchan := make(chan int, 32)
	i := 0
	for ; i < 2; i++ {
		go sendNum(numchan)
		go sendMessage(msgchan)
	}
	for {
		select {
		case msg := <-msgchan:
			fmt.Printf("Worked on %s\n", msg)
		case x := <-numchan:
			fmt.Printf("I got %d \n", x)
		}
	}
}
