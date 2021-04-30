package main

import (
	"fmt"
	"unsafe"
)

//整数类型的使用
func main() {
	var i int = 2
	fmt.Println(i)

	//有符号
	//int8的范围为：-128~127
	//其他的int16,int32,int64类推
	var j int8 = 127
	fmt.Println(j)

	//无符号
	//uint8的范围为：0~255
	//其他的uint16,uint32,uint64类推

	//扩展
	//int ,uint,rune,byte的使用
	var a int = 64
	fmt.Println(a)
	var b uint = 1
	fmt.Println(b)
	var c rune = 12
	fmt.Println(c)
	var d byte = 255
	fmt.Println(d)

	//整型的使用细节
	var n1 = 100 //n1是什么类型
	//查看变量类型
	//fmt.Printf() 可以用于做格式化输出
	fmt.Printf("%T \n", n1)

	//如何在程序查看某个变量的字节大小和数据类型
	var n2 int = 10
	//unsafe.Sizeof(n2)  是unsafe包的一个函数，可以返回n2变量占用的字节数
	fmt.Printf("n2的数据类型 %T n2占用的字节数是 %d", n2, unsafe.Sizeof(n2))
}
