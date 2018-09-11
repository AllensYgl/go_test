package baidu

type Inter struct {
	Value string
}

func (i Inter) Get(url string) string {
	return i.Value + url
}
