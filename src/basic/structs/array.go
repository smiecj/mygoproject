package main

import "fmt"

func main() {

	// 初始化一个空数组
	var intArr [10] int
	for index := 0; index < len(intArr); index++ {
		intArr[index] = index + 1
	}

	// 求和
	sum := 0
	for _, element := range intArr {
		sum += element
	}

	fmt.Printf("数组求和的结果为: %d\n", sum)

	// 直接初始化非空数组
	var notEmptyStrArr = [4] string {"锄禾日当午", "汗滴禾下土", "谁知盘中餐", "粒粒皆辛苦"}

	// 获取子数组，注意子数组所占用的对象引用和父数组是一样的，这点和Java类似，Java是返回一个SubList内部类
	// 本质上也是属于同一块内存区域，直接操作是有风险的
	var subStrArr [] string = notEmptyStrArr[1:]
	subStrArr[1] = "床前明月光"

	for index, element := range notEmptyStrArr {
		fmt.Printf("字符串数组，第%d个元素的值为 %s\n", index, element)
	}

	// 另外和Java的基本数组类型一样，go 的数组也不能直接对数组本身进行修改操作
}