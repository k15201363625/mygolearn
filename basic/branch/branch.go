package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func testif() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%s\n", contents)
	}
	//可以在条件判断中完成赋值操作并且用;隔开 实现判断
	if contents2, err2 := ioutil.ReadFile(filename); err2 == nil {
		fmt.Println(contents2)
	} else {
		fmt.Println(err2)
	}
	//contents2 只能在条件范围内生效
}

//switch 不需要表达式 不需要break 可以有多个表达式
func testswitch(a, b int, op string) int {
	var res int
	switch op { //switch 可以没有表达式 直接判断
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*", "/": //多个条件
		res = 0
	default:
		panic("unsupported op:" + op)
	}
	return res
}

//go 中没有while 而for可以实现while所有能力
func testfor() int { //for 不需要括号 可以省略三个条件
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}
func converttobin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// for 可以只有终止条件 逐行读取文件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	//返回bool类型
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	testif()
	res1 := testswitch(1, 2, "+")
	fmt.Println(res1)
	fmt.Println(testswitch(1, 2, "*"))
	sum := testfor()
	fmt.Println(sum)
	fmt.Println(converttobin(11))

}
