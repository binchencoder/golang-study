package main

import "fmt"

func main() {
	c := make(chan int, 1)

	select {
	case c <- 10: // c中放入了10, 因为chan的buffer为1
		fmt.Println("write 10")
	default:
		fmt.Println("default1")
	}

	select {
	case c <- 11:
		fmt.Println("write 10")
	default: // c中只有10, 没有11
		fmt.Println("default2")
	}

	select {
	case v, ok := <-c:
		// 读出来一个, v=10, ok=true
		fmt.Printf("v: %v, ok: %v \n", v, ok)
	default:
		fmt.Println("default3")
	}

	select {
	case v, ok := <-c:
		fmt.Printf("v: %v, ok: %v \n", v, ok)
	default: // 没有可读的，走这个分支
		fmt.Println("default4")
	}
}
