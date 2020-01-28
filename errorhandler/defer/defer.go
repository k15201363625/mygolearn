package main

import (
	"bufio"
	"errors"
	"fmt"
	//别名引入
	fib "mygolearn/mook/errorhandler/defer/fib"
	"os"
)

/*
错误处理与资源回收
*/
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	//return后面的语句根本不会执行 即便是defer
	//defer通过栈方式执行
	defer fmt.Println(4)
	fmt.Println(5)
}
func tryDefer2(num int) {

	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic(errors.New("error"))
	//panic后面的语句根本不会执行 即便是defer
	//defer通过栈方式执行

	if num == 0 {
		panic(errors.New("defer test error"))
	}
	defer fmt.Println(4)
	fmt.Println(5)

}

//总结：
/*只会执行到没有退出时的defer语句
对于退出后的defer语句不会执行
defer语句是否执行由执行时动态的退出位置决定*/

func writeFile(filename string) {
	file, err := os.Create(filename)
	//错误处理
	//文件必须不存在才不会报错
	//file,err := os.OpenFile(filename,os.O_EXCL|os.O_CREATE, 0666)
	//正确的错误处理方式
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			//未知错误
			panic(err)
		} else {
			//open fib.txt file exists
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
	}
	//close
	defer file.Close()

	//加速读写
	//bufio.NewScanner NewWriter
	writer := bufio.NewWriter(file)
	//由于只是bufferio 所以需要flush 否则没有内容
	//defer 想到什么写什么
	defer writer.Flush()

	//fprintln 实现向指定的writer写入内容
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

//defer在调用的时候动态计算参数 而不是最后才计算
//很重要
func tryDefer3() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}

//使用defer场景--资源管理
//open/close lock/unlock
//printheader printfooter
func main() {
	//tryDefer()
	//tryDefer2()
	//tryDefer2(1)
	//fmt.Println("----------------")
	//tryDefer2(0)

	//writeFile("fib.txt")
	//tryDefer3()

	//错误处理
	//处理掉已知类型 对于未知类型panic
	//能处理的尽量处理
	writeFile("fib.txt")
}
