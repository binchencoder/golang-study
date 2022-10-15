package main

import (
	"fmt"
	"sync"
	"time"
)

var MAX_THREADS = 2

// http://www.grantcss.com/blog/2019/06/21/golang-handle-fixed-threads/

/**
写过 golang 多线程的应该都知道，多线程之间可以通过管道通信协作，作为通信的数据。
我们可以通过 sync.WaitGroup 和管道配合来达到控制多线程的线程数。

sync.WaitGroup 是一个线程控制组, 它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine 执行完成。

它一共有 3 个方法，Add()，Done()，Wait()。

Add - 设置需要阻塞的线程数 Done - 代表单线程已经执行完毕 Wait - 等待所有线程执行完毕的阻塞方法

通过它就可以控制多少个消费者同时在处理了，而不是生成无限多个消费者。
**/
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
