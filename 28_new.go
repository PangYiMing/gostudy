package main
import "fmt"

func main ()  {
	// var a int = 10

	var p *int
	p = new(int) // 返回地址给p
	// 变量内存
	*p = 666

	fmt.Printf("p = %p, *p = %d \n", p,*p);
	q := new(int) // 自动推导
	*q = 77
	fmt.Printf("q = %p, *q = %d \n", q,*q);

}