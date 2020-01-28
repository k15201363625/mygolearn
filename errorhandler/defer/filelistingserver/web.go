package main

/**
使用defer recover 自定义错误+typeassertion wrapper
实现对于错误的适当处理
不仅尽量使用error 并且对于可能的panic合理处理
还是用自定义错误更灵活满足需求

宗旨：
不用或仅在无法判断出错（意料之外）时候使用panic 其余主动使用error
真实错误信息记录在后台
只将想要展示给用户的错误信息展示给用户
不要让用户遇到由于panic没有合理处理而现实的页面错误
所以可以使用提前recover处理肯恩的panic
*/
//目前为止的web.go 中的errwrapper就是常见的错误处理套路

import (
	"github.com/gpmgo/gopm/modules/log"
	myerrhandler "mygolearn/mook/errorhandler/defer/filelistingserver/filelisting"
	"net/http"
	"os"
)

//定义一个返回error的类型  作为wrapper包装器的参数
type appHandler func(w http.ResponseWriter, r *http.Request) error

//函数式编程 输入输出都是函数
//作为函数包装器使用
func errWrapper(handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	//具体处理error 返回合适的handler
	return func(writer http.ResponseWriter,
		request *http.Request) {
		//如果还有未知错误发生 所以没有返回error
		//则需要用recover() + defer 实现控制
		//服务器出错没有退出就是使用了defer recover
		//但是如果使用服务器默认的方法 仍然会出现问题
		//但是recover也是分层次 最近的先捕捉到
		//所以在这里捕捉 可以提前与server.go中的捕捉程序 进行panic处理
		defer func() {
			if r := recover(); r != nil {
				log.Warn("panic: %v", r)
				//显示用户所希望看到的 而不是崩溃的页面
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			//自己先后台记录下问题
			log.Warn("Error handling request:%s\nerr: %T,%s",
				request.URL.Path, err, err)

			//处理自定义错误
			//需要单独判断
			//定义 + typeassertion 实现自定义错误捕捉
			if userErr, ok := err.(MyUrlError); ok {
				//返回需要先是给用户的内容
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			//对于类型进行判断
			//如果使用最基本的匹配 由于都属于*pathError无法匹配
			//而是用判断函数可以得到详细类型
			switch {
			//os.PathError 里面真正的error

			case os.IsNotExist(err):
				//case os.ErrNotExist:
				code = http.StatusNotFound
			//case os.ErrPermission:
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			//进行错误显示：显示给用户页面
			http.Error(writer, http.StatusText(code), code)
		}

	}
}

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

//对于一些错误 我们希望能够显示给用户
//如果使用原生的错误
//无法通过 return并且接受自定义的error从而进行处理
//所以需要自定义error实现处理 方便类型比较
//用errors.New()生成的errorstring太死板
type MyUrlError interface {
	error
	Message() string
} //定义接口

func main() {
	//string , handlefunc
	http.HandleFunc("/",
		errWrapper(myerrhandler.HandleFileList))
	//使用之前指定的handlefunc 开启服务器程序
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

//目前为止的web.go 中的errwrapper就是常见的错误处理套路
