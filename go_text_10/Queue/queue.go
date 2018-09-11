package Queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	if q == nil || len(*q) == 0 {
		return true
	} else {
		return false
	}
}
