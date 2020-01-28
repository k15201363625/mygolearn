package queue

type Queue []interface{}

//使用别名方式 实现类拓展 改写

//可以存取都限定成int

//可以改掉本身的内容
/*func (q *Queue) Push(v int){
	*q = append(*q,v)
}*/
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

//func (q *Queue) Pop() int{
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head
//}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
func (q *Queue) Popint() int {
	head := (*q)[0]
	*q = (*q)[1:]
	//return head
	//type assertion
	return head.(int)
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

//对于用interface{} 是否限定
//2限定方式
//如果在参数 函数定义时已经限定了 则编写代码时候检查出问题
//如果在函数类型定义时候没有限定 而在逻辑内部用了assretion 则在运行时可能出错
//interface{} + assertion 实现动态效果
