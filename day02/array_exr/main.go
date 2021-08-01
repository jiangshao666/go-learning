package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机生成10个整数（1-100之间），倒序打印，求平均值，求最大，最小的下标，并查找是否有55

func randFind() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		val := rand.Intn(100) + 1
		arr[i] = val
	}
	fmt.Println("原始数组:", arr)
	arrLen := len(arr) - 1
	for i := 0; i < arrLen/2; i++ {
		tmp := arr[i]
		arr[i] = arr[arrLen-i]
		arr[arrLen-i] = tmp
	}
	fmt.Println("倒序数组:", arr)
	var maxIndex, minIndex, findIndex = 0, 0, -1
	var sum = 0.0
	for i := 0; i < arrLen; i++ {
		sum += float64(arr[i])
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
		if arr[i] == 55 {
			findIndex = i
		}
	}
	avgValue := sum / float64(arrLen)
	fmt.Println("平均值:", avgValue, "最大值下标:", maxIndex, "最小值下标", minIndex, "是否有55:", findIndex)
}

// 已知一个有序数组[1, 3, 5, 7, 10, 20, 23, 49], 向其中插入一个数，要求保持数组依然有序
func addToSortArr(addValue int) {
	var arr = [...]int{1, 3, 5, 7, 10, 20, 23, 49}
	var index = 0
	for i := len(arr) - 1; i > 0; i-- {
		if addValue > arr[i] {
			index = i
			break
		}
	}

	// 方式1：采用数组
	var arr2 [len(arr) + 1]int
	for i := 0; i <= index; i++ {
		arr2[i] = arr[i]
	}
	arr2[index+1] = addValue
	for i := index + 1; i < len(arr); i++ {
		arr2[i+1] = arr[i]
	}
	fmt.Println("after add:", arr2)

	// 方式2 slice
	s := arr[:index+1]
	send := arr[index+1:]
	newS := make([]int, 0, len(arr)+1)
	newS = append(newS, s...)
	newS = append(newS, addValue)
	newS = append(newS, send...)
	fmt.Println("after add:", newS)
}

// 定义一个3行4列的数组，逐个从键盘输入值，然后将四周数据设置成0
func resetAround() {
	var arr [3][4]int
	var row, col = 3, 4
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("请输入%d行第%d列数值:", i+1, j)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println(arr)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				arr[i][j] = 0
			}
		}
	}
	fmt.Println("after set:")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

// 随机生成10个数(1-100之间)， 使用冒泡排序，然后用二分查找是否有80这个数，如果有输出下标，没有则提示没有
func arrSortFind() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randV := rand.Intn(100) + 1
		arr[i] = randV
	}
	// 使用冒泡排序
	fmt.Println("before sort:", arr)
	nLen := len(arr)
	for i := 0; i < nLen-1; i++ {
		flag := false
		for j := 0; j < nLen-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("after sort:", arr)

	// 二分查找
	left, right := 0, nLen-1
	var foundIndex = -1
	for left <= right {
		middle := (left + right) / 2
		fmt.Println(middle, left, right)
		if arr[middle] == 80 {
			foundIndex = middle
			break
		} else if arr[middle] > 80 {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	if foundIndex > 0 {
		fmt.Println("找到了80，小标是:", foundIndex)
	} else {
		fmt.Println("没有找到80")
	}
}

// 给定一个数组，并给出8个整数，求出大于平均值和小于平均值的个数
func countOverAndLower() {
	var arr = [...]int{1, 3, 59, 19, 128, 27, 40, 49}
	sum := 0.0
	for _, v := range arr {
		sum += float64(v)
	}
	avgV := sum / float64(len(arr))
	numOver, numLower := 0, 0
	for _, v := range arr {
		if float64(v) >= avgV {
			numOver++
		} else if float64(v) < avgV {
			numLower++
		}
	}
	fmt.Printf("平均值: %f 大于平均值的有%d个，小于平均值的有%d个", avgV, numOver, numLower)
}

func main() {
	randFind()
	addToSortArr(4)
	resetAround()
	arrSortFind()
	countOverAndLower()
}
