package falser

type FalseRetriever struct {
	Contents     string
	FalseMessage string
}

//函数必须实现为 对象接受者 而不是pointer接受者 ？？
func (r FalseRetriever) Get(url string) string {
	if r.Contents == "" {
		return "i am false retriever"
	} else {
		return r.Contents
	}
}
