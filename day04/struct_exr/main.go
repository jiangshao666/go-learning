package main

import (
	"fmt"
	"os"
)

// 实现一个函数版本的学生管理系统，实现查询、增加、删除学生功能

var allStudents map[int64]*student

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id, name,
	}
}

func queryStudents() {
	for key, v := range allStudents {
		fmt.Printf("学号: %d 姓名: %s \n", key, v.name)
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名:")
	fmt.Scanln(&name)
	stu := newStudent(id, name)
	allStudents[id] = stu
}

func deleteStudent() {
	var id int64
	fmt.Print("请输入学号:")
	fmt.Scanln(&id)
	_, ok := allStudents[id]
	if !ok {
		fmt.Println("输入学号错误")
	}
	delete(allStudents, id)
}

func main() {
	allStudents = make(map[int64]*student, 100)

	fmt.Println("欢迎来到学生管理系统")
	for {
		fmt.Println(`
			1 查询学生列表
			2 增加学生
			3 删除学生
			4 退出
		`)
		fmt.Print("请选择你的操作:")
		var choice int = 0
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			queryStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		}
	}
}
