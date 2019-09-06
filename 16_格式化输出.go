package main
import "fmt"
func main()  {
	a := 10
	b := 'a'
	c := 2.01
	d := "asd"
	// %T 操作变量所属类型
	fmt.Printf("%T %T %T %T", a,b,c,d);

	// %d  十进制数值格式
	// %c  ascci 格式,可以按照数字转成对应的码
	// %s  字符串格式
	// %f  浮点格式
	// %v  自动匹配格式
	fmt.Printf("%T %T %T %T", a,b,c,d);

}