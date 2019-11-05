package Stack

//栈接口
type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(int)
	Pop()
	Peek(int) int
}
