package Queue

type Queue interface {
	GetSize() int
	IsEmpty() bool
	EnQueue(int)
	DeQueue() int
	GetFront(int) int
}
