package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch: //如果有数据，下面打印。但是有可能ch一直没数据
				fmt.Println("num = ", num)
				time.Sleep(3 * time.Second)
			case <-time.After(60 * time.Second): //上面的ch如果一直没数据会阻塞，那么select也会检测其他case条件，检测到后3秒超时
				fmt.Println("超时")
				quit <- true //写入
			}
		}

	}() //别忘了()

	for i := 0; i < 5; i++ {
		if i == 3 {
			close(ch)
			fmt.Println("Close ch")
		}
		fmt.Printf("ch <- %d \n", i)
		ch <- i
		// time.Sleep(time.Second)
	}

	stop := <-quit //这里暂时阻塞，直到可读
	if stop {
		fmt.Println("程序结束")
	} else {
		fmt.Println("程序未结束")
	}
}
