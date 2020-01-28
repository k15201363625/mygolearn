package myerrhandler

import "fmt"

//通过分析源码得到
/*
判断中 patherror是error的组合 同时其指针是error的实现类
所以通过patherror的类型断定 .（type）
实现typeswitch 之后取出其中真实包含的error
patherror中包含的才是明确的error类型
由于这些error底层都是一个字符串
所以即便是不同的对象 可以使用==进行值比较
从而通过== 将error类型确定
**
问题：go中对于type对象的比较 是单纯的值比较??
**
*/
type mydemo struct {
	s string
	a int
}

func DemoEqualOp() {
	a := mydemo{"hello", 5}
	b := a
	c := mydemo{"hello", 5}
	fmt.Println(a == b)
	fmt.Println(a == c)
	//true
	//true
	//没有引用等 全是值等号？？
	//比较大小默认使用值比较 而不是别的
}

//Main file has non-main package or doesn't contain main function
//func main() {
//	demoEqualOp()
//}
//main函数必须在main包中 所以需要新的包
