package main
import "fmt"

func main ()  {
	type bigint int64
	var a bigint

	
	
	fmt.Printf("请输入a type = %T", a);

	type(
		long int64
		char byte
	)
	var b long = 11
	var ch char ='a'
	fmt.Printf("请输入b = %v, ch = %c", b,ch);

}