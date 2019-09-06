package main
import "fmt"
func main()  {
	// iota 常量自动生成器，每一行加1,  同一行值一样
	// 常给常量赋值
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("hello world a= %d,b= %d, c= %d\n",a,b,c);
	// var d = iota
	// 遇到const 重置
	const e = iota
	
	fmt.Printf("hello world e = %d", e);

}