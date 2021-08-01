package main

import (
	"errors"
	"fmt"
	"os"
)

// 利用数组实现循环队列

type CircleQueue struct {
	maxSize int
	array   [5]int
	front   int
	tail    int
}

func (queue *CircleQueue) Push(val int) (err error) {
	if queue.isFull() {
		return errors.New("queue is full now")
	}
	queue.array[queue.tail] = val
	queue.tail = (queue.tail + 1) % queue.maxSize
	return err
}

func (queue *CircleQueue) Pop() (val int, err error) {
	if queue.Empty() {
		return -1, errors.New("queue is empty now")
	}
	val = queue.array[queue.front]
	queue.front = (queue.front + 1) % queue.maxSize
	return val, err
}

func (queue *CircleQueue) List() {
	if queue.Empty() {
		fmt.Println("queue is empty")
		return
	}
	tempFront := queue.front
	for i := 0; i < queue.Size(); i++ {
		fmt.Printf("[%d]=%d", i, queue.array[tempFront])
		tempFront = (tempFront + 1) % queue.maxSize
	}
	fmt.Println()
}

func (queue *CircleQueue) isFull() bool {
	return (queue.tail+1)%queue.maxSize == queue.front
}

func (queue *CircleQueue) Empty() bool {
	return queue.front == queue.tail
}

func (queue *CircleQueue) Size() int {
	return (queue.tail + queue.maxSize - queue.front) % queue.maxSize
}

func main() {

	queue := &CircleQueue{
		maxSize: 5,
		front:   0,
		tail:    0,
	}

	var val int
	for {
		fmt.Println("请输入选项:")
		fmt.Println("增加元素-1:")
		fmt.Println("弹出元素-2:")
		fmt.Println("查看队列-3:")
		fmt.Println("退出-4:")
		opt := 0
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			fmt.Println("请输入数值")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("弹出的值:", val)
			}
		case 3:
			queue.List()
		case 4:
			os.Exit(0)
		}
	}
}
