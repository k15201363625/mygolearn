package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//实现Stringer接口 指定通过%v 显示结果
//实现Reader Read
//实现Writer Write
type myreader struct {
	Context string
}

//func (m *myreader) Write(p []byte) (n int, err error) {
//	panic("implement me")
//}

func (m *myreader) String() string {
	return m.Context
}

//func (m myreader) Read(p []byte) (n int, err error) {
//	panic("implement me")
//}

func printReader(r io.Reader) {
	//不只局限在File实现类
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	a := myreader{"hello"}
	fmt.Printf("%v %s\n", a, a)

	//使用string生成的io,Reader
	s := "hi"
	fmt.Printf("%T\n", s)
	printReader(strings.NewReader(s))
}
