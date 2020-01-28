package tree

import "fmt"

//go 有封装 继承多态
//go面向接口 函数式 并发

type Node struct {
	Value       int
	Left, Right *Node
}

//go 成员函数指定 函数接受者
//接受者可以是val | pointer 都是pass_by_val
//等价于通过参数传入 只是调用语法不同
//使用pointer作为接收者 类似其他语法的this|self
//无论哪种接受者 在调用以及函数内部使用全部可以用. ,调用会智能判断传入类型
//函数接受者当做参数即可
func (node Node) PrintNode() {
	//虽然使用对象接受 仍然可能被nil调用
	//需要判断
	//但是由于是对象 所以不是nil无法比较
	//指针可以使nil 对象不能是nil
	//调用这个函数 如果用指针 会有nil dereference 警告

	//if node != nil{
	fmt.Println(node.Value)
	//}
}

//等价于printnode(node treeNode)
//接受者可以是nil nil可以调用
//神奇 有时需要判断 有时可以便捷代码

//使用pointer可以修改
func (node *Node) ModifyNode(value int) {
	node.Value = value
}

//factory func 代替构造函数
//可以返回局部变量地址
//go 中 变量分配到 堆/栈
//由编译器智能决定 不是定死的 并且会智能gc
//不同于c++ java python
//所以可以返回局部变量地址
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}
