package main

import "fmt"

func main() {

	//方式1: 多变量声明，一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	//一次性声明多个变量方式,声明多个并赋值(类型是自动推断)
	var n4, n5, n6 = "李四", "男", 12
	fmt.Println(n4, n5, n6)

	//一次性声明多个变量方式,声明多个并赋值(类型是自动推断)
	n7, n8, n9 := "王五", "男", 23
	fmt.Println(n7, n8, n9)
}
