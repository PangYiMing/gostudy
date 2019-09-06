package main
import "fmt"
func test() (a, b, c int)  {
	return 10,11,12
}

func main()  {
	fmt.Println("hello world\n");
	
	i, j := 10, 20
	// i, j = j, i
	fmt.Printf("hello world i= %d,j= %d \n",i,j);
	var a, b, c int
	a, b, c = test()
	// _ 匿名变量，丢弃数据不处理，配合返回值
	// a, b, _ = test()
	a, b = b, a
	fmt.Printf("hello world a= %d,b= %d, c= %d",a,b,c);

}