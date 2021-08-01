package main

import "fmt"

// append 为切片动态追加元素
func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v, len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s1[3] = "广州" //错误，为导致编译错误，索引越界
	// 注意： append的返回值必须指定一个变量来接管，如果是原来的切片则相当于扩容
	s1 = append(s1, "广州")
	// 如果是新的切片接管则相当于重新定义了一个切片 类似 s2 := append(s1, "广州")

	// 但是本质上都是底层新建了一个新的数组，并且拷贝了原有数组的元素到新的数组

	fmt.Printf("s1=%v, len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v, len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"西安", "长沙", "苏州"}
	s1 = append(s1, ss...) //append切片, ...表示拆开切片
	fmt.Printf("s1=%v, len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	// 切片没有专用的删除元素的方法， 可以通过append来实现
	// 切片删除元素并不影响底层数组的元素，因此容量不会变, 但是数组对应索引的值会被修改
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	//要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a)

}
