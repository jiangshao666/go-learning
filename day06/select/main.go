package main

import (
	"fmt"
	"math/rand"
	"time"
)

// select 多路复用

func f1() {
	var ch chan int = make(chan int, 1)
	// select如果有多个满足条件则随机选择一个，如果只有一个满条件则选择满足条件的分支
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func f2() {
	var ch chan int = make(chan int, 1)
	var ch2 chan int = make(chan int, 1)
	var ch3 chan int = make(chan int, 1)

	go func() {
		for i := 0; i < 10; i += 3 {
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
			ch <- i
		}
	}()

	go func() {
		for i := 1; i < 10; i += 3 {
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
			ch2 <- i
		}
	}()

	go func() {
		for i := 2; i < 10; i += 3 {
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
			ch3 <- i
		}
	}()

	// 一次读取case对应管道的数据，如果不匹配则逐个向后匹配，最后走默认的default

	for {
		select {
		case v := <-ch:
			fmt.Println("ch1", v)
		case v2 := <-ch2:
			fmt.Println("ch2", v2)
		case v3 := <-ch3:
			fmt.Println("ch3", v3)
		default:
			// 	fmt.Println("没有可以读取的数据了")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	//f1()
	f2()
}
