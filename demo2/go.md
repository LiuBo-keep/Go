### 整数类型
有符号
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200503094352910.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzQzMDcyMzk5,size_16,color_FFFFFF,t_70)

无符号
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200503101714371.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzQzMDcyMzk5,size_16,color_FFFFFF,t_70)

补充
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200503102141918.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzQzMDcyMzk5,size_16,color_FFFFFF,t_70)

整型的使用细节
1.Golang各个整数类型分：有符号和无符号，int,uint的大小和系统有关

2.Golang的整型默认声明为int型

3.如何在程序查看某个变量的字节大小和数据类型
   
    //查看数据类型
    fmt.Printf() 可以用于做格式化输出
    fmt.Printf("%T",n1)//使用%加大写的T
    
    //查看字节大小与数据类型
    //如何在程序查看某个变量的字节大小和数据类型
    var n2  int =10
    //unsafe.Sizeof(n2)  是unsafe包的一个函数，可以返回n2变量占用的字节数
    fmt.Printf("n2的数据类型 %T n2占用的字节数是 %d",n2,unsafe.Sizeof(n2) )
    

4.Golang程序中整型变量的使用时，遵守保小不保大的原则，即：在保证程序
正常运行下，尽量使用占用空间小的数据类型。

5.bit：计算机中最小的存储单位。

6.byte：计算机中基本存储单元