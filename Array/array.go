package Array

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
)

//var (
//	data.data []int //这里定义一个切片以下说的数组其实是切片
//)

type Arr struct {
	Data []int
}

var MyArr = Arr{
	Data: []int{},
}

//获取数组元素个数
func GetSize() int {
	return len(MyArr.Data)
}

//获取数组容量
func GetCap() int {
	return cap(MyArr.Data)
}

//判断数组是否为空
func IsEmpty() bool {
	return len(MyArr.Data) == 0
}

//最后添加元素
func AddLast(e int) {
	MyArr.Data = append(MyArr.Data, e)
	fmt.Println(MyArr.Data)
}

func AddFirst(e int) {
	s := []int{e} //定义一个只有e的切片
	MyArr.Data = append(s, MyArr.Data...)
	fmt.Println(MyArr.Data)
}

//向指定位置添加元素
func Add(index, e int) {
	//todo 因为使用切片，内存是可以自动扩容的，所以不需要判断长度
	//if len(Data) == cap(Data) {
	//	log.Fatal("Add failed. Array is full.")
	//}
	//
	//这里index >= len(Data)是因为index主要用于下标，若数组长度=index，则该方法有问题，应该直接AddLast方法
	if index < 0 || index >= len(MyArr.Data) {
		log.Fatal("Add failed. index < 0 or index >= size.")
	}

	MyArr.Data = append(MyArr.Data[:index], append([]int{e}, MyArr.Data[index:]...)...)

}

//获得index索引位置的元素
func Get(index int) {
	if index < 0 || index >= len(MyArr.Data) {
		log.Fatal("Get failed. index < 0 or index >= size.")
	}
	fmt.Println(MyArr.Data[index])
}

//改变某一位置的值
func Set(index, e int) []int {
	if index < 0 || index >= len(MyArr.Data) {
		log.Fatal("Get failed. index < 0 or index >= size.")
	}

	return append(MyArr.Data[:index], append([]int{e}, MyArr.Data[index+1:]...)...)
}

//获得数组中的数据
func GetRes() []int {
	fmt.Println("数组中的数据为：", MyArr.Data)
	return MyArr.Data
}

//查找数组中是否有某个元素
func Contains(e int) bool {
	for i := 0; i < len(MyArr.Data); i++ {
		if MyArr.Data[i] == e {
			return false
		}
	}
	return false
}

//查找e在数组中的索引，不存在返回-1
func Find(e int) int {
	for i := 0; i < len(MyArr.Data); i++ {
		if MyArr.Data[i] == e {
			return i
		}
	}
	return -1
}

//查找e在数组中的全部索引，不存在返回-1
func FindAll(e int) []int {
	var indexs []int
	for i := 0; i < len(MyArr.Data); i++ {
		if MyArr.Data[i] == e {
			indexs = append(indexs, i)
		}
	}
	return indexs
}

//删除指定位置的元素
func Del(index int) bool {
	if index < 0 || index >= len(MyArr.Data) {
		log.Fatal("Get failed. index < 0 or index >= size.")
	}
	MyArr.Data = append(MyArr.Data[:index], MyArr.Data[index+1:]...)
	return true
}

//元素若存在则删除
func RemoveElement(e int) bool {
	index := Find(e)
	if index != -1 {
		Del(index)
		return true
	}
	return false
}

//删除元素全部元素
//todo 这里好像有bug
func RemoveAllElement(e int) {
	for i, v := range MyArr.Data {
		if v == e {
			MyArr.Data = append(MyArr.Data[:i], MyArr.Data[i+1:]...)
		}
	}
	fmt.Println("last Data", MyArr.Data)
}
