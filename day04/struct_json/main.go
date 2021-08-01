package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json
// go语言中标识符如果是小写字母开头，则表示只有本包可以见，如果大写字母开头则表示对外可见
type person struct {
	name string
	age  int
}

type personEx struct {
	Name string `json:"name" db:"name" ini:"name"` //表示对该字段做名称转换
	Age  int    `json:"age"`
}

func main() {
	p := person{
		"zhangsan", 10,
	}
	b, err := json.Marshal(p) //夸包的情况json包获取不到person中的字段，因为person中的字段都是小写字母开头，只能在本包可见
	if err != nil {
		fmt.Println("json serialize failed")
		return
	}
	fmt.Println(string(b))

	p1 := personEx{
		"zhangsan", 10,
	}

	b, err = json.Marshal(p1)
	if err != nil {
		fmt.Println("json serialize failed")
		return
	}
	fmt.Println(string(b))

	// 反学列话
	s := `{"name":"lisi", "age":20}`
	var p2 personEx
	json.Unmarshal([]byte(s), &p2) //注意不能直接用p2而是用&p2，因为函数传参都是值传递
	fmt.Printf("%#v\n", p2)
}
