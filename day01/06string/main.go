package main

import (
	"fmt"
	"strings"
)

// 字符串
// GO语言中双引号包裹的是字符串, 单引号包裹的是字符
// 转义符号: \n \r \t \' \"
func main() {
	path := "D:\\golang\\src\\github.com\\jiangshao666\\day01"

	fmt.Println(path)

	s2 := "i'm OK"
	fmt.Println(s2)
	//多行字符串
	s3 := `标题
		副标题
		开头`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))
	// 字符串拼接 + 或者 fmt.Sprintf
	name := "lilei married "
	marry := "hanmeimei"
	ss := name + marry
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, marry)
	fmt.Println(ss1)
	// Split 分割

	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	// 包含 contains
	fmt.Println(strings.Contains(name, "lilei"))
	// 前缀
	fmt.Println(strings.HasPrefix(name, "lilei"))
	// 后缀
	fmt.Println(strings.HasSuffix(name, "lilei"))

	// 字串查找
	fmt.Println(strings.Index(name, "i"))
	fmt.Println(strings.LastIndex(name, "i"))

	// 拼接
	fmt.Println(strings.Join(strings.Split(path, "\\"), "+"))

	traversalString()

	changeString()
}

// go中有两种字符 一种是byte类型，用于表示ASCII码范围的字符 实际uint8
// 另一种是rune类型的用于表示中文、日文等符合字符, 实际是一个int32

func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c)", s[i], s[i])
	}

	fmt.Println()

	for _, r := range s { //rune
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
}

// 修改字符串
// go中字符串本身不可以修改，如果需要修改需要转换成byte或者rune[]才可以
func changeString() {
	s := "白萝卜"
	//s[0] = '红'

	s1 := []rune(s) //把字符串转换成rune切片
	s1[0] = '红'

	fmt.Println(string(s1)) //把rune切片强制转换成string

}
