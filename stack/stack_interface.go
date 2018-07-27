package gost

type Stack interface {
	Peek() interface{}
	Pop() interface{}
	Push(data interface{})
	Size() int
}
