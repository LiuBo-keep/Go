package main

import "fmt"

func main() {

	//第一种：指定变量类型，声明后不若赋值，使用默认值
	//int 的默认值是 0
	var i int
	fmt.Println(i)

	//方式2：根据值自行判断变量类型
	var name = "李四"
	fmt.Println(name)

	//方式3：省略var，注意 :=左侧的变量不应该是已声明过的，否则会导致编译错误
	//下面的方式等价于：var name string = "tom"
	//:=的:不能省略
	sex := "男"
	fmt.Println(sex)

	//总结：
	/*
	  方式1更适合声明全局变量
	  方式3更适合声明局部变量
	*/
}
