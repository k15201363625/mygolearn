package main

import "fmt"

func testmap() {
	m := map[string]string{
		"a": "b",
		"c": "d",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	//empty map
	m2 := make(map[int]string)
	//nil
	var m3 map[string]int

	fmt.Println(m2, m3)

	//以上是三种初始化

	//访问元素
	//实际上返回1/2 个元素 对于不存在的返回为空
	//获得的是初始值
	name := m["a"]
	fmt.Println(name)
	name = m["k"]
	fmt.Println(name)
	name, ok := m["b"]
	fmt.Println(name, ok)
	name, ok = m["c"]
	fmt.Println(name, ok)

	if nn, o := m["k"]; o != false {
		fmt.Println(nn)
	} else {
		fmt.Println("empty")
	}

	//delete
	//删除不存在的elem 不会报错 ，诶有操作
	delete(m, "b")
	fmt.Println(m)
	delete(m, "c")
	fmt.Println(m)

	//map 内部hash 所以无顺序
	//可以放到slice排序 实现有序

	fmt.Println(len(m))
	//map的key
	//除了slice map function以外的内奸类型都可以作为map的key
	//含有以上三者的struct也不能作为key
	//因为slice map funciton 认为不能比较相等？？
	//所以只要含有三者不能为key
	//map 本身没有== !=操作
	//mymap:= make(map[map[int]string] int)
	mymap2 := make(map[int]map[int]int)
	fmt.Println(mymap2)
	//map可以为value 不可以为key
	mymap3 := make(map[int][]int)
	fmt.Println(mymap3)

	//mymap4 := make(map[[]int] int)
	mymap4 := make(map[[5]int]int)
	fmt.Println(mymap4)
	//arr可以为key
	//make用于第二次赋值时 由于类型已经确定 需要相同类型
	mymap4 = make(map[[5]int]int)
	//mymap4 = make(map[[4]int] int)

}

//leetcode
func lengthOfNonRepeatingSubStr(s string) int {
	//by hash
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		//每次需要更新hashtable
		lastOccurred[ch] = i
	}
	return maxLength
}

//rune == go char
func testrune() {
	s := "慕课网" //每个中文字幕需要3
	//utf-8使用可变长
	//utf-16 2bits
	//utf-8 1bit for e ,3bits for ch
	fmt.Println(len(s))
	// %s 字符串形式表示 %X|x 16进制表示
	fmt.Println([]byte(s))
	fmt.Printf("%s", []byte(s))
	fmt.Printf("%x", []byte(s))

}

//string默认 如果使用len显示的是字节长度
//使用[]byte(s) 得到字节格式 还是原始内存空间 只是新的理解
//使用[]rune(s) 得到新的对象 每个原始char占据4bits 不是原是内存空间
//使用k,v := range s 得到的是自适应的 bytenum,char 是智能变动的
//如 1 s,2 a,5 木,8 哈,11 k,12 g
//使用utf8下的函数可以实现真实字符数目统计 以及字符序列解码
//使用strings. 下的函数可以实现基本的字符串操作

//上一个题目只需要使用[]rune(s) 实现迭代就可以处理中文问题

func main() {
	//testmap()
	//len1 := lengthOfNonRepeatingSubStr("")
	//len2 := lengthOfNonRepeatingSubStr("bbbb")
	//len3 := lengthOfNonRepeatingSubStr("abcbbbcdfg")
	//fmt.Println(len1,len2,len3)
	//但是由于使用byte处理 不能处理中文 只能处理英文字符
	//想要处理 需要rune类型

	testrune()

}
