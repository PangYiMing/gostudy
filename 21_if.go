package main
import "fmt"

func main ()  {
	s:= "wa"
	i:= 1
	if s == "wa" {
		fmt.Println("作用");
	}
	// if  支持一个初始化语句，初始化语句和判断条件以封号分割 ,声明的i不能在下面执行
	if d := 2; d == 0{
		fmt.Println("作用2");
	} else if i== 2{
		fmt.Println("作用3");

	}
	fmt.Println("4>3?", 4>3);

	switch i {
		case 1:
			fmt.Printf("i == %d \n", i);
			// break  //默认包含，可以不写
			// fallthrough // 不挑出，继续执行
		case 2:
			fmt.Printf("i == %d \n", i);
			fallthrough
		case 3:
			fmt.Printf("i == %d \n", i);
		default:
			fmt.Printf("default i == %d \n", i);
	}

	// if  支持一个初始化语句，初始化语句和判断条件以封号分割 ,声明的i不能在下面执行
	switch num:= 4; num {
	case 1, 4:
		fmt.Printf("num == %d \n", num);
		// break  //默认包含，可以不写
		// fallthrough // 不挑出，继续执行
	case 2:
		fmt.Printf("num == %d \n", num);
		fallthrough
	case 3:
		fmt.Printf("num == %d \n", num);
}
newN:= 1;
switch  {
case newN<10:
	fmt.Printf("newN < 10 , newN == %d \n", newN);
	// break  //默认包含，可以不写
	// fallthrough // 不挑出，继续执行
case newN==10:
	fmt.Printf("newN < 10 , newN == %d \n", newN);
case newN>10:
	fmt.Printf("newN < 10 , newN == %d \n", newN);
}
}