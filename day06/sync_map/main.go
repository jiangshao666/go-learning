package main

import (
	"fmt"
	"strconv"
	"sync"
)

// go语言中内置的map不是线程安全的, 要使用线程安全的map可以使用sync.Map
var m1 map[string]int

func innerMap() {
	m1 = make(map[string]int)
	var wg sync.WaitGroup
	for i := 0; i < 19; i++ {
		wg.Add(1)
		go func(key string, v int) {
			m1[key] = v
			fmt.Printf("get k=%s v=%d\n", key, m1[key])
			wg.Done()
		}(strconv.Itoa(i), i)
	}
	wg.Wait()
}

var m2 sync.Map

func syncMap() {
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(key string, v int) {
			m2.Store(key, v)
			val, _ := m2.Load(key)
			fmt.Printf("get k=%s v=%d\n", key, val)
			wg.Done()
		}(strconv.Itoa(i), i)
	}
	wg.Wait()
}

func main() {
	//innerMap()
	syncMap()
}
