package main

import "fmt"

type Teacher struct {
	id int16
	name string
	subject string
}

func main() {
	tXiaoMing := Teacher{1, "Mr.Li", "science"}
	tXiaoMingPointer := &tXiaoMing
	fmt.Println(tXiaoMing)
	fmt.Printf("输出老师的信息: %+v\n", tXiaoMing)
	fmt.Printf("输出老师的姓名: %s\n", tXiaoMing.name)
	// 指针也可以直接通过.来访问属性
	fmt.Printf("输出老师的课程: %d\n", tXiaoMingPointer.subject)
}