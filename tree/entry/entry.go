package main

import (
	"fmt"
	"mygolearn/mook/tree"
)

//继承 拓展：通过组合/别名 两种方式实现
//组合
type myTreeNode struct {
	node *tree.Node
}

//组合实现后序遍历
//注意类型转换
//过程中需要创建大量的组合后的类型的变量
func (myNode *myTreeNode) postOrderTranverse() {
	if myNode == nil || myNode.node == nil {
		return
	}
	//不支持直接使用初始化对象的指针作为地址
	//需要生成新的对象
	left := myTreeNode{node: myNode.node.Left}
	//left = myTreeNode(myNode.node.Left)
	//无法通过类型转换实现 由于是组合而不是传统的继承
	//组合中只能通过国川建新变量方式 不能通过类型转换
	//但是使用别名 可以实现类型转换
	left.postOrderTranverse()
	right := myTreeNode{node: myNode.node.Right}
	right.postOrderTranverse()
	myNode.node.PrintNode()
}

//封装 + 包
//通过命名实现封装
//main函数作为入口必须在main package中
//一个包只能位于一个单独的文件夹下
//一个struct的成员函数 可以在一个包中不同的文件中
//大写开头表示包内公有 小写开头表示私有 都是针对包而言
//包查找：GOPATH + GOROOT 下的src目录

//gopath
//推荐所有个人项目+第三方在gopath下

func main() {
	//使用new创建的不是nil
	//new创建的是默认初始化好的
	//直接var声明pointer的是nil
	//直接var声明的对象是默认初始化好的

	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	//root.Right.Right = new(tree.Node{Value:5})
	root.Right.Right = tree.CreateTreeNode(2)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(root, nodes)

	//创建对象
	//使用classname{field:val}指定字段值实现
	//使用new(classname) new中只有类型 没有初始化列表
	//创建数组可以省去每个前面的classname

	//指针问题
	//.表示指针引用属性 不需要->

	root.Tranverse()

	//nil 调用函数
	//默认初始化好的
	//proot := new(tree.Node)
	//var proot tree.Node

	//nil
	//proot1 := new(*tree.Node)
	//new 返回的是对象的指针 这里是指针的指针

	var proot *tree.Node
	var proot2 tree.Node
	//无法将nil 与 对象类型比较
	//不同于c++ java中null定义
	fmt.Println(proot2) //{0 <nil> <nil>}

	//if proot2 == nil{
	//	fmt.Println("proot2 is nil")
	//}else{
	//	fmt.Println(proot2)
	//}

	if proot == nil {
		fmt.Println("proot is nil")
	}
	//if proot1 == nil{
	//	fmt.Println("proot1 is nil")
	//}
	fmt.Println("hi", proot)
	//fmt.Println(proot1)

	//proot1.printnode()
	//proot1.tranverse()
	//需要用nil判断保护
	//proot.printnode()报错 nil dereference
	proot.Tranverse()

	mynode := myTreeNode{&root}
	mynode.postOrderTranverse()

	fmt.Println("------------")
	root.Tranverse2()
	//实现数据统计
	//内部可以使用的局部变量
	count := 0
	sum := 0
	root.TranverseFunc(func(node *tree.Node) {
		//关键点执行的函数
		count++
		sum += node.Value
	})
	fmt.Printf("count:%d\nsum:%d\n", count, sum)

}
