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

	// 另外和Java的基本数组类型一样，go 的数组也不能直接对数组本身进行修改操作，比如append、remove等

	// 尝试插入超过数组大小的值
	dynamicArr := make([]string, 0)
	dynamicArr = append(dynamicArr, "尝试插入新元素")
	fmt.Printf("动态数组的信息: 大小: %d, 容量: %d\n", len(dynamicArr), cap(dynamicArr))
	for index, ele := range dynamicArr {
		fmt.Printf("动态数组的第%d个元素: %s\n", index, ele)
	}

	// 测试append 之后的数组是否和原来的数组占用的内存空间一样
	initDynamicArr := make([]string, 5)
	initDynamicArr[0] = "第1个元素"

	firstAppendArr := append(initDynamicArr, "第一次扩容")
	secondAppendArr := append(firstAppendArr, "第二次扩容")
	firstAppendArr[1] = "第2个元素"

	fmt.Printf("第一次扩容之后的数组: len=%d, cap=%d, %v\n", len(firstAppendArr), cap(firstAppendArr), firstAppendArr)
	fmt.Printf("第一次扩容之后的数组: len=%d, cap=%d, %v\n", len(secondAppendArr), cap(secondAppendArr), secondAppendArr)

	// 二维数组
	xyArr := [][]uint {
		{1,2,3,4},
		{5,6,7,8},
	}

	// 遍历二维数组
	for _, ele := range xyArr {
		fmt.Printf("当前二维数组中的元素: %v\n", ele)
	}

	// 当前看到: map
    // https://tour.golang.org/moretypes/19
}