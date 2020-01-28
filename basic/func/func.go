package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 名字 参数 返回类型
//支持多个返回 还可以给返回值命名
func div(a, b int) (q, r int) {
	return a / b, a % b
	//q = a/b
	//r = a%b
	//return 这种方法仅仅使用于简单的函数
}

//对于不想要的返回 用_
//常用语：val, error双返回

//函数式编程
func apply(op func(int, int) int, a, b int) int {
	//获取函数名称
	//通过反射获取指针 之后通过运行时方法从指针获取函数指针获取名字
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName)
	return op(a, b)
}
func mypow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//不支持函数重载
//func mypow(a,b string) string{
//	return 5
//}

//go 函数不支持默认参数 不支持参数列表,可选参数 还不支持函数重载
//仅仅支持可变参数
func sumArgs(values ...int) int {
	sum := 0
	for i := range values {
		//range 数组 得到下表索引列 并且可以用for迭代
		sum += values[i]
	}
	return sum
}

func main() {
	q, r := div(5, 6)
	fmt.Println(q, r)
	//main.mypow
	fmt.Println(apply(mypow, 5, 6))
	//使用匿名函数方式
	//main.main.func1
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
}
