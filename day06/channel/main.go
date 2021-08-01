package main

import (
	"fmt"
	"sync"
)

// channel 语法 x := <-chan 或者 chan <- value

// channel close是可选项，如果不手动close，则当内容读取完成后也会自动回收
// close的channel具有如下特点:
/*
	1) 不可以再写入，如果写入会panic
	2) 可以继续读取内容，直到读取完成
	3) 如果对已经读取完成的channle继续读取，会读取到指定类型的0值
	4) 对已经close的channel再次close是panic
*/

// 无缓冲区的chan, 如果没有对应接收逻辑会导致chan卡死
func nobuffChan() {
	var ch chan int
	fmt.Println(ch)
	ch = make(chan int)

	go func() {
		x := <-ch
		fmt.Println("读取到的值", x)
	}()

	ch <- 10
	fmt.Println("发送完成")
	close(ch)
}

// 带缓冲区的chan,再缓冲区写满之前可以一直往里面写入，但是如果写满以后依然没有读取，则会卡死
func bufChan() {
	var ch chan int = make(chan int, 1)
	ch <- 1 //可以正常放入
	x := <-ch
	fmt.Println(x)
	ch <- 2 //可以正常放入
	ch <- 3 //panic，缓冲区满
	close(ch)
}

var wg sync.WaitGroup
var once sync.Once
var counter = 0

// 独立的goroutine对channel循环写入
// chan<- 表示一个单项的只写channel
func writeChan(ch chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

// 独立的goroutine从ch1循环读取，然后对读取的值平方再放入ch2
func readWriteChan(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			counter++
			break
		}
		fmt.Println("read i", x)
		ch2 <- x * x
	}
	if counter == 2 { //最后一个goroutine负责关闭ch2
		once.Do(func() { close(ch2) })
	}
}

func main() {
	//nobuffChan()
	//bufChan()
	var ch1 chan int = make(chan int, 100)
	var ch2 chan int = make(chan int, 100)
	wg.Add(3)
	go writeChan(ch1)
	go readWriteChan(ch1, ch2)
	go readWriteChan(ch1, ch2)

	wg.Wait()
	// 从ch2读取最终的结果
	for x := range ch2 {
		fmt.Println(x)
	}

}
