package main

import "fmt"

// 环形链表
type Node struct {
	data int
	next *Node
}

type CircleLink struct {
	head *Node
}

func (cl *CircleLink) AddNode(val int) {

	if cl.head.next == nil { //插入第一个节点
		cl.head.data = val
		cl.head.next = cl.head
	} else {
		node := &Node{data: val}
		// 定位到尾节点
		temp := cl.head
		for temp.next != cl.head {
			temp = temp.next
		}
		temp.next = node
		node.next = cl.head
	}
}

// 按照值删除,用于值不重复的情况
func (cl *CircleLink) DelNode(val int) {
	pCur := cl.head
	if pCur.next == nil {
		fmt.Println("link is empty")
		return
	}
	for pCur.next.data != val && pCur.next != cl.head {
		pCur = pCur.next
	}
	if pCur.next == cl.head {
		cl.head = pCur.next.next
	}
	pCur.next = pCur.next.next

}

// 按照位置删除
func (cl *CircleLink) DelAt(index int) {
	if index < 1 {
		fmt.Println("index out of range")
		return
	}
	if index == 1 { //删除头结点
		tail := cl.head.next
		if tail == cl.head { //只有一个节点
			cl.head.next = nil
			return
		}
		for tail.next != cl.head {
			tail = tail.next
		}
		tail.next = cl.head.next
		cl.head = cl.head.next

	} else {
		temp := cl.head
		flag := false
		for i := 2; temp.next != cl.head; i++ {
			if i == index {
				flag = true
				break
			}
			temp = temp.next
		}
		if flag {
			temp.next = temp.next.next
		} else {
			fmt.Println("index out of range")
		}
	}
}

func (cl *CircleLink) List() {
	temp := cl.head
	for {
		fmt.Printf("%d -> ", temp.data)
		if temp.next == cl.head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

// 循环链表的应用
// 约瑟夫环：1-n个人围城一个圈，从第K个人开始计数，数到第M个则退出，然后从后续的人重新计数，循环进行知道所有人退出，打印退出顺序

func josephuCircle(cl *CircleLink, k int, m int) {
	pCur := cl.head
	if pCur.next == cl.head {
		fmt.Println("只有一个人", cl.head.data)
		return
	}
	// 移动到初始定位点
	for i := 0; i < k-1; i++ {
		pCur = pCur.next
	}

	for {
		if pCur.next == pCur {
			fmt.Println("最后退出:", pCur.data)
			return
		}
		for i := 1; i < m-1; i++ {
			pCur = pCur.next
		}
		fmt.Println("退出:", pCur.next.data)
		pCur.next = pCur.next.next //删除退出节点
		pCur = pCur.next           //重置重新计数起点
	}
}

func main() {
	cl := &CircleLink{
		head: &Node{},
	}
	cl.AddNode(1)
	cl.AddNode(2)
	cl.AddNode(3)
	cl.AddNode(4)
	cl.AddNode(5)

	//cl.DelAt(6)
	//cl.DelNode(1)
	cl.List()

	josephuCircle(cl, 2, 3)

	//cl.List()
}
