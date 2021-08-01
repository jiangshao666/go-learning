package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁，针对读多写少的情况

var (
	x      = 0
	lock   sync.Mutex
	rwlock sync.RWMutex
	wg     sync.WaitGroup
)

func read() {
	defer wg.Done()
	//lock.Lock()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	//lock.Lock()
	rwlock.Lock()
	x++
	time.Sleep(time.Millisecond * 5)
	//lock.Unlock()
	rwlock.Unlock()
}

func main() {
	start := time.Now()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go write()
	}
	time.Sleep(time.Millisecond * 200)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
