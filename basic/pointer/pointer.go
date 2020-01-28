package main

import "fmt"

//go pointer
//不能进行指针运算
//书写类似c++
func testpointer() {
	var a int = 5
	var p *int = &a
	*p = 3
	fmt.Println(a)
}

//pass_by_val
//pass_by_ref:c++之外所有自定义类型
//go只有值传递唯一方式
//所以go参数传递与指针配合
//指针的值传递==引用传递

func swap(a, b *int) {
	//a,b = b,a错误
	*a, *b = *b, *a
}
func swap2(a, b int) (int, int) {
	return b, a
}
func main() {
	testpointer()
	a := 5
	b := 6
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println(swap2(5, 6))
}
