package main

import (
	"fmt"
	"strings"
)

// 测试：设计一个方法，输入一串字符串，输出字符串中的每一个单词出现的次数
func getWordAndTimes(article string) map[string]int {
	wordCountMap := make(map[string]int)
	wordArr := strings.Fields(article)
	for _, word := range wordArr {
		wordCountMap[word] = wordCountMap[word] + 1
	}
	return wordCountMap
}

// 测试：闭包功能，设计斐波那契数列
// 0 1 1 2 3 5 8 13 21
func fabonacciFunc() func() int {
	// 这里定义的其实是初始值
	lastlastNum, lastNum := 0, 1
	return func() int {
		currentNum := lastlastNum
		lastlastNum = lastNum
		lastNum = currentNum + lastNum
		return currentNum
	}
}

func main() {
	mathTeacher := Teacher{123, "Mr.Liang", "Math"}

	mymap := make(map[string]Teacher)

	mymap["math"] = mathTeacher

	// 遍历map
	fmt.Println("开始遍历动态map！")
	for key, value := range mymap {
		fmt.Printf("当前元素的key 为: %s，值为: %v\n", key, value)
	}

	// 声明的时候直接初始化map
	fmt.Println("开始遍历初始化map！")
	myInitMap := map[string]Teacher{
		"chinese": {1, "Mrs.zhang", "chinese"},
		"english": {2, "Mr.huang", "english"},
	}
	for key, value := range myInitMap {
		fmt.Printf("当前元素的key 为: %s，值为: %v\n", key, value)
	}

	// 删除元素，并确认
	delete(myInitMap, "chinese")
	chineseTeacher, hasChineseTeacher := myInitMap["chinese"]
	fmt.Printf("当前语文老师是否存在: %t, 值为: %v\n", hasChineseTeacher, chineseTeacher)

	// 测试获取一篇文章当中所有单词出现的次数
	wordCountMap := getWordAndTimes("This is my hello world go project!!!")
	fmt.Println("输出单词和出现次数的统计信息: ")
	for key, value := range wordCountMap {
		fmt.Printf("当前元素的key 为: %s，值为: %v\n", key, value)
	}

	// 测试闭包编写斐波那契数列
	funcOfFab := fabonacciFunc()
	for index := 0; index < 10; index++ {
		fmt.Printf("当前第%d个斐波那契数为: %d\n", index + 1, funcOfFab())
	}

	// 下一次看到：方法 https://tour.golang.org/methods/1
}