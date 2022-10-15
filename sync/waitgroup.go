package main

import (
	"fmt"
	"sync"
	"time"
)

var MAX_THREADS = 2

func main() {
	jobs := make(chan int, 10)

	var wg sync.WaitGroup
	// 设置需要多少个线程阻塞
	wg.Add(MAX_THREADS)
	fmt.Println(time.Now())

	// 根据线程数控制启动多少个消费者线程
	for n := 0; n < MAX_THREADS; n++ {
		go worker(jobs, &wg)
	}

	// 生产者
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	// 等待所有线程执行完毕的阻塞方法
	wg.Wait()
	fmt.Println(time.Now())
}

// 消费者
func worker(jobs <-chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Println(job)
		time.Sleep(1 * time.Second)
	}
	// 消费完毕则调用 Done，减少需要等待的线程
	wg.Done()
}
