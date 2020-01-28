package myerrhandler

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//实现自己的error
type userError string

func (err userError) Message() string {
	return string(err)
}

func (err userError) Error() string {
	return err.Message()
}

const prefix string = "/list/"

func HandleFileList(w http.ResponseWriter, r *http.Request) error {
	//内部处理逻辑
	//进行提前错误处理
	if strings.Index(r.URL.Path, prefix) != 0 {
		//return errors.New(fmt.Sprintf("path should start with %s",prefix))
		//返回自定义的error
		//注意 由于使用的是别名 初始化一个别名类型 string->userError
		myerr := userError("path should start with" + prefix)
		return myerr
	}
	filepath := r.URL.Path[len(prefix):]
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	//bufio用于包装writer reader
	//ioutil 用于使用writer reader实现读取写入操作
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(all)
	return nil
}
