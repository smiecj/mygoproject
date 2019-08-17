package main

import "fmt"

type Student struct {
	id int16
	name string
	sex string
	grade int
	classId int16
}

func main() {
	sXiaoMing := Student{1, "xiaoming", "boy", 1, 1}
	sXiaoMingPointer := &sXiaoMing
	fmt.Println(sXiaoMing)
	fmt.Printf("输出小明的信息: %+v\n", sXiaoMing)
	fmt.Printf("输出小明的姓名: %s\n", sXiaoMing.name)
	// 指针也可以直接通过.来访问属性
	fmt.Printf("输出小明的年级: %d\n", sXiaoMingPointer.grade)
}