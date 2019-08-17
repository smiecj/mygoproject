package loop

import (
	"fmt"
	"math"
)

func loopTestBasic(upperBound int) int {
	sum := 0
	for i := 0; i < upperBound; i++ {
		sum += i
	}
	return sum
}

func loopTestWhile() int {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	return sum
}

// 测试找到最接近x 的平方数
func testSqrt(x float64) float64 {
	if x <= 2 {
		return x
	}
	// 首先找到最接近的整数
	var beginNum int64 = 2
	var lastNum int64 = 1
	for beginNum * beginNum < int64(x) {
		lastNum = beginNum
		beginNum = beginNum * beginNum
	}
	beginNum = lastNum

	// 然后进行小数上的运算操作，精度自定
	precision := 100
	beginFloatNum := float64(beginNum)
	beginCalFloatNum := 0.5

	for index := 0; index < precision; index++ {
		if currentNum := beginFloatNum + beginCalFloatNum; currentNum * currentNum < x {
			beginFloatNum += beginCalFloatNum
		} else {
		}
		beginCalFloatNum = beginCalFloatNum / 2
	}
	return beginFloatNum
}

func loop() {
	fmt.Printf("测试循环加: %d\n", loopTestBasic(20))
	fmt.Printf("测试while循环加: %d\n", loopTestWhile())

	var sqrtNum float64 = 3
	fmt.Printf("sqrt精度对比: %f 的完全平方数计算: %f, %f", sqrtNum, testSqrt(sqrtNum), math.Sqrt(sqrtNum))

	// 0817: 继续学习Defer，一个比较特殊的用法，延迟执行。参考: https://blog.golang.org/defer-panic-and-recover
	for index := 1; index <= 10; index++ {
		defer fmt.Println(index)
	}
}