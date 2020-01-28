package myerrhandler

import (
	"errors"
	"fmt"
	"os"
)

//上面发现了patherror作为error的组合type 可以转换成error类型
//本身是因为不仅是组合 也实现了接口

type demo struct {
	val int
}
type demo2 struct {
	value demo
}

func helpfunc(v int) {
	fmt.Println(v)
}
func helpfunc2(v demo) {
	fmt.Println(v)
}
func helpfunc3(v demo2) {
	fmt.Println(v)
}
func demoCombinationTypeConvert() {

	//test := demo{5}
	//test2 := demo2{demo{4}}

	//var a int = int(test)  no
	//helpfunc(test)  no
	//组合类型无法相互转换 即便使用指针也不可以
	//helpfunc2(test2) no
	//helpfunc3(test) no
	//helpfunc2(&test2) no
	//helpfunc3(&test) no

	testerr := os.PathError{"a", "b", errors.New("test")}
	//os.IsPermission(testerr)
	os.IsPermission(&testerr)
	//使用指针就可以实现转换 神奇
	//本身是因为不仅是组合 也实现了接口

}
