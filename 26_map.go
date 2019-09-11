package main
import "fmt"
func main()  {
	// 只有len，没有cap
	var ma  map[int]string
	fmt.Println(" ma = ", ma);
	fmt.Println(" len(ma) = ", len(ma));

	ma2 := make(map[int]string)
	fmt.Println(" ma2 = ", ma2);
	fmt.Println(" len(ma2) = ", len(ma2));

	// 指定长度相当与指定容量，但是里面一个数据没有
	ma3 := make(map[int]string, 10)
	fmt.Println(" ma3 = ", ma3);
	fmt.Println(" len(ma3) = ", len(ma3));
	ma4 := make(map[int]string)
	ma4[1] = "s333s"
	ma4[0] = "ss3s"
	ma4[5] = "ss1"
	fmt.Println(" ma4 = ", ma4);
	fmt.Println(" len(ma4) = ", len(ma4));

	// 初始化
	// key唯一

	ma5 := map[int]string{1:"ss", 2:"go"}
	fmt.Println(" ma5 = ", ma5);
	fmt.Println(" len(ma5) = ", len(ma5));
}