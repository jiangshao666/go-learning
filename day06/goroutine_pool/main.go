package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d start job %d \n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d end job %d \n", id, j)

		result <- j * 2
	}
}

func main() {
	var jobs chan int = make(chan int, 100)
	var result chan int = make(chan int, 100)

	// 启动三个worker
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, result)
	}

	// 生成一些任务
	for j := 0; j < 10; j++ {
		jobs <- j
	}

	close(jobs)

	for j := 0; j < 10; j++ {
		r := <-result
		fmt.Println("result", r)
	}
}
