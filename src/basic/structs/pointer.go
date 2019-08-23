package main

import (
	"basic/model"
	"fmt"
)

func addOne(num *int) {
	*num = *num + 1
}

func addStudentGrade(student *model.Student) {
	student.Grade++
}

func main() {
	num := 1
	addOne(&num)
	fmt.Printf("1+1=%d\n", num)

	// 测试通过指针，修改结构体内部属性
	xiaoming := model.Student{1, "xiaoming", 1}
	addStudentGrade(&xiaoming)
	fmt.Printf("学生%s的信息为: %v\n", xiaoming.Name, xiaoming)
}