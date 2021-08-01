package main

import (
	"fmt"
	"sync"
)

// 锁，保证对公共资源的并发修改的安全性

var x int = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		mutex.Lock()
		x++
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
