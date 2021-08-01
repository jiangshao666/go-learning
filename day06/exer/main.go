package main

import (
	"fmt"
	"sync"
)

/*
启动1个goroutine往一个chan中放入1-2000数
启动8个goroutine，完成1-2000个数累加 1-n的和，然后存入一个resChan中，最后遍历输出resChan的结果
*/

var once = sync.Once{}

type calcResult struct {
	index int
	sum   int
}

func putChan(ch chan<- int) {
	for i := 1; i <= 2000; i++ {
		ch <- i
	}
	close(ch)
}

func calcChan(ch <-chan int, resCh chan<- *calcResult, wg *sync.WaitGroup) {
	for {
		v, ok := <-ch
		if ok {
			sum := 0
			for i := 1; i <= v; i++ {
				sum += i
			}
			calcRes := &calcResult{v, sum}
			resCh <- calcRes
		} else {
			wg.Done()
			break
		}
	}
	wg.Wait()

	once.Do(func() {
		close(resCh)
	})
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 200)
	resCh := make(chan *calcResult, 200)
	go putChan(ch)
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go calcChan(ch, resCh, &wg)
	}
	for {
		v, ok := <-resCh
		if !ok {
			break
		} else {
			fmt.Printf("res[%d]=%d\n", v.index, v.sum)
		}
	}
}
