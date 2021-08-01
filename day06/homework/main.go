package main

import "fmt"

/*
	启动一个goroutine，往一个channel放入1-8000的数, 然后启动4个gouroutine分别从channel读取数，判断是否是素数
	将素数放入一个channel，最后遍历素数的channel打印结果

*/

func putNum(ch chan<- int) {
	for i := 1; i <= 8000; i++ {
		ch <- i
	}
	close(ch)
}

func calcPrime(ch <-chan int, primeChan chan<- int, exitChan chan<- bool) {
	for {
		val, ok := <-ch
		if !ok {
			break
		} else {
			isPrime := true
			for j := 2; j <= val/2; j++ {
				if val%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primeChan <- val
			}
		}
	}
	fmt.Println("primeChan finish")
	exitChan <- true
}

func main() {
	var putChan = make(chan int, 64)
	var primeChan = make(chan int, 64)
	var exitChan = make(chan bool, 4)

	go putNum(putChan)

	for i := 0; i < 4; i++ {
		go calcPrime(putChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		ret, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(ret)
	}
}
