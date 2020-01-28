package tree

func (node *Node) Tranverse() {
	if node == nil {
		return
	}
	node.Left.Tranverse()
	node.PrintNode()
	node.Right.Tranverse()
}

//函数式编程实例3
//func作为参数
//类似java中传入实现某个接口的类
//java中通过匿名函数/lambda表达式实现某个接口中的函数
//从而实现间接地将函数作为参数传入 表面上是传入了实现某个接口的类的对象
//并且由于匿名函数存在，保证可以使用上文中局部变量定义函数
//go中也可以通过匿名函数作为参数
//同时实现使用局部变量定义函数
/*
这是对于变量的动态检测
不同于普通函数定义 需要有确定的变量
在匿名函数作用下 变量可以使外部定义的局部变量 通过动态监测获得
*/

//e.g.3
//参数可以是函数或者函数的指针??
func (node *Node) TranverseFunc(myfun func(*Node)) {
	if node == nil {
		return
	}
	//中序遍历逻辑
	node.Left.TranverseFunc(myfun)
	myfun(node)
	node.Right.TranverseFunc(myfun)
}

func (node *Node) Tranverse2() {
	if node == nil {
		return
	}
	//传入匿名函数
	node.TranverseFunc(func(node *Node) {
		node.PrintNode()
	})
}
