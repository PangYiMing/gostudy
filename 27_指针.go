package main
import "fmt"

func main ()  {
	var a int = 10

	// 变量内存
	fmt.Println("a =", a);
	// 变量地址
	fmt.Printf("&a = %p\n", &a);


	// 定义是特殊的声明
	// 指针类型
	var p *int 
	p = &a
	fmt.Printf("p = %p，&a = %p\n", p,&a);
	// *p 操作的不是p的内存，是指p地址指向的内存
	*p = 666


	// 不要操作没有合法指向的内存
	// p = nil
	// *p = 666 // 会报错

	
}