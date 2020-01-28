package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//fibonacci
/*func fibonacci() func() int{
	a,b := 0,1
	return func() int{
		a,b = b,a+b
		return a
	}
}*/
//如果不用intGet 会无法识别 别名局限性
func Fibonacci() intGet {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//对于函数 使用函数实现接口
//函数是一等公民
//可以作为 参数 返回类型 普通类型 接口实现者
type intGet func() int

//用这个类型作为实现者 体现接口灵活性
func (g intGet) Read(
	p []byte) (n int, err error) {
	//fib func本身类似于一个可以迭代的文件
	//所以实现为reader
	next := g()
	//终止条件
	if next > 10000 {
		return 0, io.EOF
		//error是interface
		//errors.New()返回了一个实现了error接口的对象
		//想要实现error接口需要实现函数Error() string
	}
	//使用现成的reader作为代理实现
	s := fmt.Sprintf("%d\n", next)
	//TODO: incorrect if p is too small
	//通过缓存数字 检查p长度 实现
	return strings.NewReader(s).Read(p)
}

func printContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
