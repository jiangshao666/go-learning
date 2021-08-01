package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map
func main() {
	// map的定义 var 变量名 map[Key类型]value类型
	// map是引用类型，所以必须要初始化才可以使用

	var m1 map[string]int         //只是定义了类型和变量，但是没有分配内存
	m1 = make(map[string]int, 10) //第二个参数指定容量
	m1["aa"] = 1                  //赋值
	m1["bb"] = 2

	fmt.Println(m1)
	fmt.Println(m1["bb"]) //取值

	val, ok := m1["cc"]
	if !ok {
		fmt.Println("没有对应key")
	} else {
		fmt.Println(val)
	}

	// 如果不用ok的方式判断，默认获取的不存在的key的指是类型的0值
	fmt.Println(m1["cc"])

	// map的遍历 for range
	for key, val := range m1 {
		fmt.Printf("%v %v\n", key, val)
	}

	// 只取key
	for key := range m1 {
		fmt.Println(key)
	}

	//只取value
	for _, value := range m1 {
		fmt.Println(value)
	}

	// 删除元素 delete, 如果删除的元素不存在则不做任何操作
	delete(m1, "aa")
	fmt.Println(m1)

	//顺序遍历map的方式，先遍历map把key取出来放入一个slice，然后对slice做排序，最后遍历slice依次去除对应map中的value

	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	// 去除所有的key存入切片
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	//对切片排序
	sort.Strings(keys)
	// 按照顺序遍历key和获取对应的map中的value
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}
