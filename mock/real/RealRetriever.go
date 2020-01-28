package realr

import (
	"fmt"
	/*	"net/http"
		"net/http/httputil"*/)

type Realr struct {
	Contents  string
	UserAgent string
}

//通过指针可以实现修改
func (r *Realr) Post(url string, params map[string]string) {
	r.Contents = url
	fmt.Println(r.Contents)
	fmt.Println(params)
}

func (r *Realr) Get(url string) string {
	/*	resp,err := http.Get(url)
		if err != nil{
			panic(err)
		}
		res,err := httputil.DumpResponse(resp,true)
		if err != nil{
			panic(err)
		}
		resp.Body.Close()
		//[]byte -> string
		r.Contents = string(res)
		return r.Contents*/

	return r.Contents
}
