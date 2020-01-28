package main

import (
	"fmt"
	"math"
)

/*
表格驱动测试 vs 传统测试
1 数据逻辑分离
2 错误信息明确 传统测试使用assert 错误信息混乱
3 不会因为一个错误中指 可以有多个错误

go 很容易实现表格驱动测试
*/
func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}
func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
func main() {
	triangle()
}
