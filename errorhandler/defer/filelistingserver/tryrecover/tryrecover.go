package main

import "fmt"

func tryRecover() {
	defer func() {
		//recover自动接收panic参数中收到的量
		r := recover()
		//interface{} 可以转换成任意类型
		//使用typeassertion
		//返回的是任意类型 需要转换
		if err, ok := r.(error); ok {
			fmt.Println("error recovered:", err)
		} else {
			panic(fmt.Sprintf("unknowned error%T %v\n", r, r))
		}
	}() //defer必须加上函数调用

	//a := 5
	//b := 0
	//fmt.Println(a/b)
	panic("hello")
	/*
	   panic: hello [recovered]
	   panic: unknowned errorstring <nil>
	   会有两次的panic记录
	*/

}
func main() {
	tryRecover()
}
