package main

import "fmt"

//数据结构：
//数组 切片 容器

//数组定义
func testarr() {
	//三种定义方法
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	//支持多维数组
	arr4 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var arr5 [3][4]int
	fmt.Println(arr1, arr2, arr3, arr4, arr5)

	//遍历数组
	//range三种用法 可以同时获取i v
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
}

//传入数组
//数组是是值传递 所以值类型数组不会改变
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//即便使用了指向数组的指针，同样在内部调用与普通数组变量相同
//不需要(*arr)[i]
func printArray2(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//go中通常不直接使用数组/数组的指针
//通常使用切片 []int
//slice 切片
func testslice() {
	arr := [...]int{1, 2, 3, 5, 4, 8, 6}
	s := arr[2:6] //前开后闭
	fmt.Println(s)
	fmt.Println(arr[2:])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])
}

//切片实质上存储的是视图 执行会关联修改
func passbyslice(arr []int) {
	fmt.Println(arr)
	arr[0] = 5
}

//支持reslice操作

//slice 深入 {ptr,len,cap}
//可以向后拓展 不能向前拓展
//向后拓展不能超过cap
//arr := [5]{1,2,3,4,5}
//s1 = arr[1:2]
//s2 = arr[3:5]可以
func testslice2() {
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:2]
	s2 := arr[3:5] //可以
	fmt.Println(s1, s2)
	fmt.Println(cap(s1), cap(s2))
}

//append 作用到slice
//使用append
//如果没有超过 cap 则还是view
//如果超过了 cap 则生成新的arr作为底层arr
//因为不再是view 不会改变原来的arr
//原来的数组如果没有人用 会垃圾回收

func testappend() {
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:2]
	s2 := arr[3:5] //可以
	s3 := append(s1, 10)
	fmt.Println(arr)
	s4 := append(s2, 15)
	fmt.Println(arr)
	fmt.Println(s3, s4)
}

func printSlice(s []int) {
	fmt.Println(len(s), cap(s))
}

//懂爱扩大容量 *2
func createslice() {
	var s []int
	//s == nil
	printSlice(s)
	for i := 1; i <= 10; i++ {
		s = append(s, i)
	}
	printSlice(s)

	s1 := []int{2, 4, 5, 6}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

}
func sliceops() {
	//append
	//copy
	s2 := make([]int, 16)
	s1 := []int{1, 2, 3, 4}
	copy(s2, s1) //不需要接受返回
	fmt.Println(s2)

	//delete internal elem
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)
	//通过切片... 实现展开变成变长参数

	//delete front or tail
	front := s2[0]
	s2 = s2[1:]
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail)
	fmt.Println(s2)

}
func main() {
	//testarr()
	//printArray([5]int{1,2,3,4,5})
	//printArray2(&[...]int{5,6,7,8,9})
	//testslice()
	//demoarr := []int{1,2,3,4,5}
	//passbyslice(demoarr[:])
	//fmt.Println(demoarr)

	//testslice2()
	//testappend()

	//createslice()
	sliceops()
}
