package demo2

import "fmt"

func main()  {

	//golang的变量声明使用方式1
	//第一种：指定变量类型，声明后不若赋值，使用默认值
	//int 的默认值是 0
	var i int
	fmt.Println("i=",i)

	//方式2：根据值自行判断变量类型
	var mun =1.25
	fmt.Println("mun=",mun)

	//方式3：省略var，注意 :=左侧的变量不应该是已声明过的，否则会导致编译错误
	//下面的方式等价于：var name string = "tom"
	//:=的:不能省略
	name := "tom"
	fmt.Println("name=",name)

	//总结：
	/*
	  方式1更适合声明全局变量
	  方式3更适合声明局部变量
	 */
}
