package main

import (
	"fmt"
	falser "mygolearn/mook/mock/false"
	realr "mygolearn/mook/mock/real"
)

/*duck typing
不关注内部结构 关注外部功能

严格意义上：要求动态绑定（go不算）

<tempelate R> //c++
<R extends Receiver> // java
//python没有
string download(R r){
return r.get("www.baidu.com")
}

python中duck typing
通过注释说明接口 运行时才知道 需要注释

c++中的duck typing
通过摸板支持 <tempelate R>编译时才知道 还需要注释

java中duck typing
通过模板支持 <R extends Receiver>
编写代码时候就知道 需要实现Receiver接口的类才能传入
严格来说不算duck typing
如果需要实现多个接口 无法实现有效限制

目标：
有java类型检查
具有python c++ 灵活的ducktyping

go中接口：
传统语言由实现者实现接口 使用者有一定定义能力
go中由使用者定义
go中接口实现者只需要实现所需要的函数 不需要指明实现哪个接口 具有灵活性
go中接口只有实现了相应方法type才能作为接口的变量 具有java的检查功能
虽然不是严格的动态绑定 但是也具有了类型检查+灵活性
*/

type Retriever interface {
	Get(url string) string
}

func download(retriever Retriever) string {
	return retriever.Get("http://www.mook.com")
}

//使用r.(type)得到类型 方法一
func inspect(r Retriever) {
	fmt.Printf("%T,%v\n", r, r)
	switch v := r.(type) { //检查类型
	case *realr.Realr:
		fmt.Println("realretriever:", v)
	case falser.FalseRetriever:
		fmt.Println("falseretriever:", v)
	}
}

//使用type assertion
//类似动态绑定 运行时候判断
//
func dynamicTypeAssertion(r Retriever) {
	//myr := r.(*realr.Realr)
	//fmt.Println((*myr).UserAgent)
	if myr, ok := r.(*realr.Realr); ok { //两个返回的判断
		fmt.Println((*myr).UserAgent)
	} else {

		otherr := r.(falser.FalseRetriever)
		//当assertion 失败 产生nil对象 并且这个nil对象不能被再次赋予新的类型
		//*realr.Realr,<nil>
		//由于类型已经确定为assertion的类型
		fmt.Printf("%T,%v\n", myr, myr)
		fmt.Println(otherr.FalseMessage)
	}
}

//接口组合
type Poster interface {
	Post(string, map[string]string)
}

type RetrieverPoster interface {
	Retriever
	Poster
	//还可以有别的方法的要求 但是这种模式很常见
}

func session(rr RetrieverPoster, url string) string {
	rr.Post(url, map[string]string{
		"name": "gmm",
	})
	return rr.Get(url)
}
func main() {
	var r Retriever
	var rr Retriever
	//r = realr.Realr{}
	//需要用&r 与接受参数匹配
	r = &realr.Realr{"HELLO", "aaa"}
	//由于是私有的 所以不能直接初始化 **
	rr = falser.FalseRetriever{"hello", "bbb"}

	//fmt.Printf("%T,%v\n%T,%v\n",r,r,rr,rr)
	//fmt.Println(download(rr))
	//fmt.Println(download(r))     	fmt.Println("--------------------------")

	//获得类型
	inspect(r)
	inspect(rr)

	fmt.Println("--------------------------")
	dynamicTypeAssertion(r)
	fmt.Println("--------------------------")
	dynamicTypeAssertion(rr)

	fmt.Println("--------------------------")
	myr := realr.Realr{}
	session(&myr, "hello")
	inspect(&myr)
	//效果 全空
	dynamicTypeAssertion(&myr)
	//总结
	//1 实现了接口的组合的类型可以作为单个接口的实现者身份
	//2 实现接口的组合的对象不能assertion为单独接口的实现类型
	//3 实现接口组合的类 只能冒充单个接口的实现类(给接口赋值)
	//	但是无法真实转换过去
	//	**并且一旦转换成单一接口的实现类 不能再assertion回来**
	// e.g.
	//  myr := &realr.Realr{}
	//	var a Retriever = myr  (ok)
	//  myr.(*realr.Realr) (no) 失败 为空
}

//接口内部可以是val|指针类型
//保存了一个类型+一个val
//类型 + 实现者的指针/值
//接口变量由于自带指针 所以作为参数通常不需要传递接口的指针
//接口中含有的
//如果是指针：只能被指针赋值
//如果是值：可以被接口或者值赋值

//type switch
//type assertion
//interface{}表示任意类型
//详细使用在queue.go
