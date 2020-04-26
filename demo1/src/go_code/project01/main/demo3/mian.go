package demo3

import "fmt"

//定义全局变量
var school = "aaa"
var techer = "bbb"
//上面的方式也可以改成一次性声明
var (
	aa = 300
	bb = 400
	namea = 12
)

func mian()  {
	//方式4: 多变量声明，一次性声明多个变量

	var n1, n2, n3 int

	fmt.Println(n1,n2,n3)

	//一次性声明多个变量方式,声明多个并赋值(类型是自动推断)
	var n4, name,n5 =100,"tom",888
	fmt.Println(n4,name,n5)

	//一次性声明多个变量方式,声明多个并赋值(类型是自动推断)
	n6,age,n7:=100,"12",888
	fmt.Println(n6,age,n7)

	//使用全局变量
	fmt.Println(school,techer)
	//使用全局变量方式2(当局部变量与全局变量一样，遵循就近原则)
	fmt.Println(aa,bb,namea)
}
