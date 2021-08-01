package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 常见排序算法 冒泡、选择、插入、快速, 下述默认都按照升序排列

const ARR_LEN = 100000

func BubbleSort(arr *[ARR_LEN]int) {
	nLen := len(arr)
	nStart := time.Now().Unix()
	for i := 0; i < nLen-1; i++ {
		for j := 0; j < nLen-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	nEnd := time.Now().Unix()
	fmt.Println("Bubble cost=", nEnd-nStart)
}

func SelectSort(arr *[ARR_LEN]int) {
	nLen := len(arr)
	nStart := time.Now().Unix()
	for i := 0; i < nLen-1; i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < nLen; j++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	nEnd := time.Now().Unix()
	fmt.Println("SelectSort cost=", nEnd-nStart)
}

func InsertSort(arr *[ARR_LEN]int) {
	nLen := len(arr)
	nStart := time.Now().Unix()
	for i := 1; i < nLen; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	nEnd := time.Now().Unix()
	fmt.Println("InsertSort cost=", nEnd-nStart)
}

// 快排填坑法
func partion(arr *[ARR_LEN]int, low int, high int) int {
	pivotVal := arr[low]
	for low < high {
		for high > low && arr[high] >= pivotVal {
			high--
		}
		if low < high {
			arr[low] = arr[high]
		}
		for low < high && arr[low] <= pivotVal {
			low++
		}
		if low < high {
			arr[high] = arr[low]
		}
	}
	arr[low] = pivotVal
	return low
}

// 快排交换法
func partion2(arr *[ARR_LEN]int, low int, high int) int {
	pivotVal := arr[low]
	l := low
	r := high
	for l < r {
		for r > l && arr[r] > pivotVal {
			r--
		}

		for l < r && arr[l] <= pivotVal {
			l++
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	arr[l], arr[low] = arr[low], arr[l]
	return l
}

// 顺序遍历法
func partion3(arr *[ARR_LEN]int, low int, high int) int {
	pivotVal := arr[high]
	index := low
	for i := low; i < high; i++ {
		if arr[i] < pivotVal {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[index], arr[high] = arr[high], arr[index]
	return index
}

func QuickSort(arr *[ARR_LEN]int, left int, right int) {
	if left < right {
		// i := partion(arr, left, right)
		// i := partion2(arr, left, right)
		i := partion3(arr, left, right)

		QuickSort(arr, left, i-1)
		QuickSort(arr, i+1, right)
	}
}

func main() {
	//arr := [12]int{0, 7, 11, 9, 6, 6, 4, 2, 1, 9, 0, 1}
	arr := [ARR_LEN]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < ARR_LEN; i++ {
		arr[i] = rand.Intn(ARR_LEN)
	}
	// fmt.Println("before sort", arr)
	selectArr := [ARR_LEN]int{}
	for i := 0; i < ARR_LEN; i++ {
		selectArr[i] = arr[i]
	}

	insertArr := [ARR_LEN]int{}
	for i := 0; i < ARR_LEN; i++ {
		insertArr[i] = arr[i]
	}

	quikArr := [ARR_LEN]int{}
	for i := 0; i < ARR_LEN; i++ {
		quikArr[i] = arr[i]
	}

	// fmt.Println("before sort arr", arr)
	// fmt.Println("before sort selectArr", arr)
	// fmt.Println("before sort insertArr", arr)
	// fmt.Println("before sort quikArr", arr)
	BubbleSort(&arr)

	// fmt.Println("after bubble sort", arr)

	SelectSort(&selectArr)
	// fmt.Println("after select sort", selectArr)
	for i := 0; i < ARR_LEN; i++ {
		if selectArr[i] != arr[i] {
			fmt.Println("select sort != bubble sort")
			break
		}
	}

	InsertSort(&insertArr)
	// fmt.Println("after insert sort", insertArr)
	for i := 0; i < ARR_LEN; i++ {
		if insertArr[i] != arr[i] {
			fmt.Println("insert sort != bubble sort")
			break
		}
	}
	nStart := time.Now().Unix()
	QuickSort(&quikArr, 0, ARR_LEN-1)
	nEnd := time.Now().Unix()
	// fmt.Println("after quick sort", quikArr)
	fmt.Println("QuickSort cost=", nEnd-nStart)
	for i := 0; i < ARR_LEN; i++ {
		if quikArr[i] != arr[i] {
			fmt.Println("quiksort != bubble sort")
			break
		}
	}
}
