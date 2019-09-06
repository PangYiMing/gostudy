package main
import "fmt"

func main ()  {
	var flag bool
	flag = true
	fmt.Printf("flag %v \n", flag)
	fmt.Printf("flag %t \n", flag)

	// bool 类型  int 类型 不能互转，为不兼容类型
	// fmt.Printf("flag %d \n", flag)

	var ch  byte
	ch = 'a'
	

	var t int
	t = int(ch)
	fmt.Printf("flag %v \n", t)

}