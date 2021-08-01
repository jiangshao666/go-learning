package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// sync waitgroup
// 监听其他gorutine执行完成
// gorutine退出的条件: 1) gorutine的任务函数执行完成后退出 2) 启动gorutine的main函数结束，则改gorutine也退出

// 随机库的用法
func random() {
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Int()    //随机一个int，范围是int的范围
	r2 := rand.Intn(10) //随机一个 [0,n)
	fmt.Println(r1, r2)
}

func f1(i int) {
	defer wg.Done() //计数减1
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	//random()

	for i := 0; i < 10; i++ {
		wg.Add(1) //计数加1
		go f1(i)
	}
	wg.Wait() //计数为0则结束
	fmt.Println("main done")
}
