package Array

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
)

var (
	data []int
)

//获取数组元素个数
func GetSize() int {
	return len(data)
}

//获取数组容量
func GetCap() int {
	return cap(data)
}

//判断数组是否为空
func IsEmpty() bool {
	return len(data) == 0
}

//最后添加元素
func AddLast(e int) {
	data = append(data, e)
	fmt.Println(data)
}

func AddFirst(e int) {
	s := []int{e} //定义一个只有e的切片
	data = append(s, data...)
	fmt.Println(data)
}

//向指定位置添加元素
func Add(index, e int) {
	//todo 因为使用切片，内存是可以自动扩容的，所以不需要判断长度
	//if len(data) == cap(data) {
	//	log.Fatal("Add failed. Array is full.")
	//}
	//

	//这里index >= len(data)是因为index主要用于下标，若数组长度=index，则该方法有问题，应该直接AddLast方法
	if index < 0 || index >= len(data) {
		log.Fatal("Add failed. index < 0 or index >= size.")
	}

	for i := len(data) - 1; i >= index; i-- {
		data = append(data[:i], append([]int{e}, data[i:]...)...)
	}

	fmt.Println(data)
}
