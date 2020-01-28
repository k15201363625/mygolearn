package main

import (
	"fmt"
	"mygolearn/mook/queue"
)

func main() {
	q := queue.Queue{1}         //ok
	qq := [...]queue.Queue{}    //no
	qqq := make(queue.Queue, 5) //ok
	qqqq := new(queue.Queue)    //ok
	//尽管是别名 仍然是不同类型 不能用其他的类赋值
	//var q5 queue.Queue = [...]int{1,2,3,4,5}
	//q = [...]int{1}
	fmt.Printf("%T\n%T\n%T\n%T\n", q, qq, qqq, qqqq)
	/*
		queue.Queue
		[0]queue.Queue ???
		queue.Queue
		*queue.Queue
		总结：
		可以将别名替换掉原类型所在位置 实现别名类型创建赋值
		但是不能用原类型初始化/赋值 别名类型
		并且[...]很玄幻 如果针对的是slice
	*/

	//由于本身是slice 别名 所以需要切片方式初始化
	//但是用法还是不同

	//不同与使用组合方式
	//使用别名方式可以实现类型转换 双向
	//使用组合必须创建新的对象 （如递归调用中）
	//[]interface {},[1]
	myq := []interface{}(q)
	fmt.Printf("%T,%v\n", myq, myq)
	//queue.Queue,[1 2 3]
	a := []interface{}{1, 2, 3}
	var mya queue.Queue = queue.Queue(a)
	fmt.Printf("%T,%v\n", mya, mya)

	q.Push(30)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	q.Pop()
	fmt.Println(q.IsEmpty())

	fmt.Println(q, qq, qqq)
}
