package Stack

//栈
//栈从尾部进入或者出去
import "../Array"

type Stacks struct {
}

func (s Stacks) GetSize() int {
	return Array.GetSize()
}

func (s Stacks) IsEmpty() bool {
	return Array.IsEmpty()
}

func (s Stacks) GetCap() int {
	return Array.GetCap()
}

func (s Stacks) Push(data int) {
	Array.AddLast(data)
}

func (s Stacks) Pop() {
	Array.RemveLast()
}

func (s Stacks) Peek(data int) int {
	return Array.GetLast()
}
