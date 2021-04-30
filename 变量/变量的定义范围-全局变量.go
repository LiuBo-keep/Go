package main

import "fmt"

// 定义全局变量

// 方式1
var school = "xxx大学"
var teacher = "xxx老师"

// 方式2
var (
	class    = "xxx班级"
	location = "xxx位置"
)

func main() {
	// 全局变量的使用
	fmt.Println("信息1=" + school + teacher)
	fmt.Println("信息2=" + class + location)
}
