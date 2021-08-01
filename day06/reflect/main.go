package main

import (
	"fmt"
	"reflect"
)

type myInt int64

type person struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

func (p person) Eat(food string) {
	fmt.Printf("person eat food:%s \n", food)
}

func (p person) Sleep() string {
	msg := "person need sleep"
	fmt.Println(msg)
	return msg
}

// 通过反射获取类型
func reflectType(a ...interface{}) {
	for _, v := range a {
		typeV := reflect.TypeOf(v)
		fmt.Printf("type:%v kind: %v \n", typeV.Name(), typeV.Kind())
	}
}

//通过反射获取值
func reflectValue(x ...interface{}) {
	for _, e := range x {
		v := reflect.ValueOf(e)
		fmt.Printf("value kind %v %T \n", v.Kind(), v)
		switch v.Kind() {
		case reflect.Int, reflect.Int64:
			fmt.Printf("type is int, value is %d\n", v.Int())
		case reflect.Float32, reflect.Float64:
			fmt.Printf("type is float, value is %f\n", v.Float())
		}
	}
}

// 通过反射设置变量的值
//想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，
//必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int32 {
		v.SetInt(200) //修改的是副本, 会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	//反射中使用专有的Elem()方法来获取指针对应的值
	if v.Elem().Kind() == reflect.Int32 {
		v.Elem().SetInt(200)
	}
}

// 结构体反射
func reflectStruct() {
	p := person{
		Name: "zhangsan",
		Age:  20,
	}
	// 属性反射
	t := reflect.TypeOf(p)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if ageField, ok := t.FieldByName("Age"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", ageField.Name, ageField.Index, ageField.Type, ageField.Tag.Get("json"))
	}

	// 方法反射
	fmt.Printf("person have methodNum:%d \n", t.NumMethod())
	v := reflect.ValueOf(p)
	for i := 0; i < v.NumMethod(); i++ {
		mName := t.Method(i).Name
		fmt.Printf("method name: %v method type: %v \n", mName, v.Method(i).Type())

		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		if mName == "Eat" {
			var args = []reflect.Value{reflect.ValueOf("apple")}
			ret := v.Method(i).Call(args)
			fmt.Println("call Eat ret:", ret)
		} else if mName == "Sleep" {
			var args = []reflect.Value{}
			ret := v.Method(i).Call(args)
			fmt.Println("call Sleep ret:", ret)
		}
	}
}

func main() {
	var a int = 0
	var b int32 = 10
	c := [...]string{"a", "b", "c"}
	var d struct {
		name string
	}
	var e myInt
	var f *float32
	reflectType(a, b, c, d, e, f)
	fmt.Println("------------------------------")
	reflectValue(a, b, c, d, e, f)
	//reflectSetValue1(b)
	reflectSetValue2(&b) //必须传递指针
	fmt.Printf("afterSetValue %d\n", b)
	reflectStruct()
}
