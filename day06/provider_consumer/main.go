package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用goroutine和channel实现一个计算int64随机数各位数字和的程序
/*
1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
2. 开启24个goroutine从jobChan中取出随机数计算各位数的个，将结果发送到resultChan
3. 主goroutine从resultChan取出结果并打印到终端输出
*/

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func provider(jobChan chan<- *job) {
	defer wg.Done()
	for {
		rValue := rand.Int63()
		newJob := &job{value: rValue}
		jobChan <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(jobChan <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		jobValue := <-jobChan
		var sum = int64(0)
		val := jobValue.value
		for val > 0 {
			sum += val % 10
			val /= 10
		}
		newResult := &result{
			job: jobValue,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	// 开启一个provider
	wg.Add(1)
	go provider(jobChan)
	// 开启24个consumer
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go consumer(jobChan, resultChan)
	}

	for r := range resultChan {
		fmt.Printf("value %d sum: %d \n", r.job.value, r.sum)
	}

	wg.Wait()
}
