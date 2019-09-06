package main
import "fmt"

func main ()  {
	var a int
	
	fmt.Println("请输入a=?");
	//  阻塞等待用户输入
		fmt.Scanf("%d =", &a);
		// fmt.Scan(&a);
	fmt.Println("a =", a);

}