package main

import (
	"errors"
	"fmt"
	"os"
)

// 双向链表

type Node struct {
	data int
	pre  *Node
	next *Node
}

type DoubleLink struct {
	head *Node
}

func (dl *DoubleLink) List() {
	if dl.head.next == nil {
		fmt.Println("link is empty")
		return
	}
	temp := dl.head
	for temp.next != nil {
		fmt.Println("node value:", temp.next.data)
		temp = temp.next
	}
}

// 逆向遍历
func (dl *DoubleLink) RevList() {
	if dl.head.next == nil {
		fmt.Println("link is empty")
		return
	}

	temp := dl.head
	for temp.next != nil {
		temp = temp.next
	}
	// 遍历到尾部了
	for temp.pre != nil {
		fmt.Println("node value:", temp.data)
		temp = temp.pre
	}
}

// 在链表尾加入node
func (dl *DoubleLink) AddNode(val int) (err error) {
	node := &Node{
		data: val,
	}
	temp := dl.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
	node.pre = temp
	return err
}

// 删除指定值的node
func (dl *DoubleLink) DelNode(val int) (err error) {
	temp := dl.head
	for temp.next != nil && temp.next.data != val {
		temp = temp.next
	}
	if temp.next == nil {
		return errors.New("not found the node of value")
	}
	temp.next = temp.next.next
	if temp.next != nil {
		temp.next.pre = temp
	}
	return err
}

func main() {

	queue := &DoubleLink{
		head: &Node{},
	}

	var val int
	for {
		fmt.Println("请输入选项:")
		fmt.Println("增加元素-1:")
		fmt.Println("弹出元素-2:")
		fmt.Println("查看队列-3:")
		fmt.Println("查看逆序队列-4:")
		fmt.Println("退出-5:")
		opt := 0
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			fmt.Println("请输入数值")
			fmt.Scanln(&val)
			err := queue.AddNode(val)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			fmt.Println("请输入数值")
			fmt.Scanln(&val)
			err := queue.DelNode(val)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			queue.List()
		case 4:
			queue.RevList()
		case 5:
			os.Exit(0)
		}
	}
}
