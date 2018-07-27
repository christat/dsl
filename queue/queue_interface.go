package gost

type Queue interface {
	Dequeue() interface{}
	Enqueue(data interface{})
	Size() int
}
