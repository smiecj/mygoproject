package main

import (
	"fmt"
	//"math/rand"
)

func add(x, y int) int {
	return x + y
}

func getPairAddAndSubResult(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return 
}

func main()  {
	// fmt.Printf("这是我的第一个测试项目！", add(1, 2))
	fmt.Printf("%s %d %s", "相加结果", add(1, 2), "\r\n")

    // 多个参数打印需要用到: 
    sum, sub := getPairAddAndSubResult(1, 2)
	fmt.Printf("%s %d %d", "同时计算两种结果", sum, sub)

	// 当前看到: https://tour.golang.org/basics/11
}
