package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//包内部变量 不是全局变量
var aa = 5

// aa:=6 不可函数外
var (
	aaa        = 5
	bbb string = "def"
)

//包内变量可以不使用
//函数内变量可以用:=定义 并且必须使用

func variable() {
	//名字在前
	var a int
	var s string
	var ss string = "abc"
	fmt.Printf("%d,%q\n", a, s)
	fmt.Println(ss)
}
func vardetection() {
	var a, b, c = 3, true, "abc"
	d, e, f := 5, false, "haha"
	fmt.Println(a, b, c, d, e, f)
}

//var type
/*
[build-in type]
bool string
int int8 16 32 64 (u) uintptr
byte rune(4字节) utf-8三个字节某些字符
float32 float64 complex64 complex128
指针
*/
func testtype() {
	// 实虚表示
	c := 3 + 4i
	fmt.Println(c, cmplx.Abs(c))
	//指数表示
	d := cmplx.Pow(math.E, 1i*math.Pi)
	e := cmplx.Exp(1i * math.Pi)
	fmt.Println(d, e)
	//python中
	//cmath.exp(1j*cmath.pi)相同结果
	//都是用float存储

}

func typeconvert() {
	//只有强制类型转换 没有隐式
	var a int
	var b, c = 3, 4
	//下取整
	a = int(math.Sqrt(float64(b*b + c*c)))
	fmt.Println(a)
}

//常量定义
const (
	a = 55
	b = 7
)

func consts() {
	//常量如果不指定类型，不会检查类型，只会直接替换
	//所以不需要类型转换
	//go中const通常不会大写
	//数值作为各种类型使用
	const a = 5

}
func enums() {
	const (
		cpp  = 0
		java = 2
	)
	fmt.Println(cpp, java)
	const (
		//go 是关键字 不能用
		golang = iota
		c
		js
	)
	fmt.Println(golang, c, js)
	//var(
	//	a = 1<<(10*iota)
	//	b
	//	cc
	//	d
	//)
	// var不能用iota iota表示自增 可以用到const赋值中 可以有共识

}
func main() {
	fmt.Println("Hello world.")
	//variable()
	//vardetection()
	//testtype()
	//typeconvert()
	enums()
}
