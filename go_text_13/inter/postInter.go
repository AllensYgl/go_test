package inter

// 接口 1
type IPost interface {
	Post() string
}

// 接口 2
type IGet interface {
	Get(value string)
}

// 接口 3   ----- 合并 1  2  功能
type Interf interface {
	IGet
	IPost
}

type AA struct {
	Value string
}

func (a AA) Get(value string) {
	a.Value = value
}

func (a AA) Post() string {
	return a.Value
}

func Sseesio(i Interf) {
	i.Get("aaa")
	i.Post()
}
