package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用单向链表实现队列

type Node struct {
	data int
	next *Node
}

type LinkQueue struct {
	head    *Node
	size    int
	maxSize int //指定队列是否有限制 -1是表示无限制
}

func (q *LinkQueue) Push(val int) (err error) {
	if q.maxSize > 0 && q.size >= q.maxSize {
		return errors.New("queue is full now")
	}
	temp := q.head
	for temp.next != nil {
		temp = temp.next
	}
	node := &Node{
		data: val,
	}
	temp.next = node
	q.size++
	return err
}

func (q *LinkQueue) Pop() (val int, err error) {
	if q.size == 0 {
		return -1, errors.New("queue is empty now")
	}
	temp := q.head
	val = temp.next.data
	q.head = temp.next
	q.size--
	return val, err
}

func (q *LinkQueue) List() {
	temp := q.head
	if q.size == 0 || temp.next == nil {
		fmt.Println("queue is empty")
	} else {
		for temp.next != nil {
			fmt.Println("node value", temp.next.data)
			temp = temp.next
		}
	}
}

func main() {

	queue := &LinkQueue{
		maxSize: -1,
		size:    0,
		head:    &Node{},
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
