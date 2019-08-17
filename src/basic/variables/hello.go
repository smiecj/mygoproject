package variables

import (
	"fmt"
	"math"
	"math/big"

	//"math/rand"
	"math/cmplx"
)

func add(x, y int) int {
	return x + y
}

func getPairAddAndSubResult(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return 
}

func testSwap(x, y int) {
	x, y = y, x
}

// 定义变量
// go 的基本类型参考网站：
var (
	myBoolean bool = false
	myMaxLong uint64 = 1 << 64 - 1
	myComplexNum complex128 = cmplx.Sqrt(12i -5)
)

const MYPICONST = 3.14159

func variables()  {
	// fmt.Printf("这是我的第一个测试项目！", add(1, 2))
	fmt.Printf("%s %d\n", "相加结果", add(1, 2))

    // 多个参数打印需要用到: 
    sum, sub := getPairAddAndSubResult(1, 2)
	fmt.Printf("%s %d %d\n", "同时计算两种结果", sum, sub)

	// 当前看到: https://tour.golang.org/basics/11

	// 打印类型信息
	fmt.Printf("类型: %T, 取值: %v\n", myBoolean, myBoolean)
	fmt.Printf("类型: %T, 取值: %v\n", myMaxLong, myMaxLong)
	fmt.Printf("类型: %T, 取值: %v\n", myComplexNum, myComplexNum)

    // 类型强转
    var x, y int = 3, 4
    var sqrtRet uint = uint(math.Sqrt(float64(x * x + y * y)))
    fmt.Printf("勾股定理: %d %d %d\n", x, y, sqrtRet)

    // 修改对象
    changeVar := 3
    fmt.Printf("一开始changeVar 的类型为: %T\n", changeVar)
    // 下一行报错：对象的类型在初始化之后就已经确定，不能自己修改
    //changeVar = 3.14
	//fmt.Printf("后续changeVar 的类型为: %T\n", changeVar)

	// 常量
	const MYTRUECONST = true
	fmt.Printf("我的两个常量: %t %f\n", MYTRUECONST, MYPICONST)

	// 大数
	// Java 里面大数要用BigInteger，但是go 不需要显式声明类型
	const myConstBigInt = 1 << 100
	fmt.Printf("我的常量大数为: %f", myConstBigInt * 1.0)

	// 用于计算的大数
	myCalBigInt := big.NewInt(math.MaxInt64)
	myCalBigInt.Mul(myCalBigInt, myCalBigInt)
	fmt.Printf("我的可计算大数为: %d\n", myCalBigInt)

	num1, num2 := 1, 2
	testSwap(num1, num2)
	fmt.Printf("测试交换元素: %d %d\n", num1, num2)
}
