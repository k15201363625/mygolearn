package main

import "fmt"

//函数式编程
//函数一等公民：变量 参数 返回值

//函数式编程 vs 函数指针

//严格的函数式编程要求
//1 不变性：不能有状态量 只能有常亮+函数
//2 单参数性 只能有一个参数

//go 中
//函数体 = 局部变量 + 自由变量
//自有变量可以连接到所有需要用到的状态量
//函数体+自由变量连成的网络 = 闭包
//go c++ java python 都支持
//java c++ 都是以前的语法可以模拟 新的语法支持
//java c++ 通过函数对象/函数指针+lambda表达式实现支持
//python 有原生支持 并且可以查看闭包中内容（自由变量)
//python中 __closure查看  python中通过定义内部函数实现 没有匿名函数

// 函数作为 返回值 参数 接口
//go : 对于自由变量访问更自然
//	   没有呢lambda表达式 但是有匿名函数

//e.g.1
//闭包中变量外部无法访问？？
func adder() func(int) int {
	sum := 0 //自由变量 状态量
	//使用匿名函数
	return func(v int) int {
		sum += v
		return sum
		//return sum += v 不能用
	}
}

//如果不使用中间变量 使用严格的函数式编程
//alias
//递归定义
type iAdder func(int) (int, iAdder)

//通过返回新函数+函数参数为状态变量实现消除状态变量
func addr2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, addr2(base + v)
	} // func(int) (int,递归定义)
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
	b := addr2(5)
	v := 0
	//小心 如果使用v,b:=b(i)
	//for 中上一次生成的新变量b会被外部b覆盖
	//从而现实没有使用的变量b??
	for i := 0; i < 10; i++ {
		v, b = b(i)
		fmt.Println(v)
	}
	//main.iAdder
	fmt.Printf("%T\n", b)
}
