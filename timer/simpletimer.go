package main

import (
	"fmt"
	"time"
)

func main() {
	//
	timer := time.NewTimer(1 * time.Second)
	<-timer.C
	fmt.Println("expired")
}
