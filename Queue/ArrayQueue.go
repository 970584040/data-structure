package Queue

import "../Array"

type Queues struct {
}

func (q Queues) GetSize() int {
	return Array.GetSize()
}

func (q Queues) IsEmpty() bool {
	return Array.IsEmpty()
}

func (q Queues) Caps() int {
	return Array.GetCap()
}

func (q Queues) EnQueue(e int) {
	Array.AddLast(e)
}

func (q Queues) DeQueue() int {
	_, data := Array.Del(0)
	return data
}

func (q Queues) GetFront(int) int {
	return Array.GetFirst()
}

func (q Queues) GetArr() []int {
	return Array.GetArr()
}
