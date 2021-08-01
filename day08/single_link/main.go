package main

import (
	"errors"
	"fmt"
)

// 单向链表
type Node struct {
	data int
	next *Node
}

type SinleLink struct {
	head *Node
}

// 头插法
func (sl *SinleLink) AddBefore(val int) {
	node := &Node{
		data: val,
	}
	node.next = sl.head.next
	sl.head.next = node
}

// 尾插法
func (sl *SinleLink) AddAfter(val int) {
	node := &Node{
		data: val,
	}
	temp := sl.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
}

// 指定位置插法, index范围从1开始
func (sl *SinleLink) AddAt(val int, index int) (err error) {
	if index <= 0 {
		return errors.New("invalid index")
	}
	pos := 1
	temp := sl.head
	for temp.next != nil && pos < index {
		pos++
		temp = temp.next
		if pos == index {
			break
		}
	}
	if pos < index {
		return errors.New("add pos out of range")
	}
	node := &Node{
		data: val,
	}
	node.next = temp.next
	temp.next = node
	return err
}

// 删除指定元素
func (sl *SinleLink) DelValue(val int) (err error) {
	temp := sl.head
	for temp.next != nil {
		if temp.next.data == val {
			break
		}
		temp = temp.next
	}
	if temp.next == nil {
		return errors.New("not found val")
	}
	temp.next = temp.next.next
	return err
}

// 删除第index个数 [index 从1计数]
func (sl *SinleLink) DelAt(index int) (err error) {
	temp := sl.head
	if index <= 0 || (index >= 0 && temp.next == nil) {
		return errors.New("index out range")
	}
	pos := 0
	for temp.next != nil && pos < index-1 { //先定位到index-1的位置
		pos++
		temp = temp.next
		fmt.Println("index and pos", pos, index, temp.next)
	}
	if pos == index-1 && temp.next == nil { //index超范围
		return errors.New("index out of range")
	}
	temp.next = temp.next.next

	return err
}

// 反转链表， 通过逐个将链表后续节点头插法实现反转
func (sl *SinleLink) Reverse() {
	if sl.head.next == nil {
		return
	}
	prev := sl.head.next
	pCur := prev.next
	for pCur != nil {
		prev.next = pCur.next
		pCur.next = sl.head.next
		sl.head.next = pCur
		pCur = prev.next
	}
}

func (sl *SinleLink) List() {
	temp := sl.head.next
	for temp != nil {
		fmt.Printf("%d ->", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	sl := &SinleLink{
		head: &Node{},
	}

	for i := 1; i <= 5; i++ {
		sl.AddBefore(i)
	}
	//sl.List()
	for i := 6; i <= 10; i++ {
		sl.AddAfter(i)
	}
	sl.List()
	// sl.AddAt(0, 10)
	// sl.List()
	// sl.DelValue(0)
	// sl.List()
	//sl.DelAt(11)
	sl.Reverse()
	sl.List()
}
