package main
import "fmt"
import "time"

func main ()  {

	for i := 0; i < 10; i++ {
		fmt.Printf("i == %d \n",i);
	}
	str := "ssssffff"
	// range 返回两个值，一个元素位置，一个元素值本身
	for i, data := range str {
		fmt.Printf("str[%d] == %c \n",i, data);
	}
	for i := range str {
		fmt.Printf("00 str[%d] == %c \n",i, str[i]);
	}
	i:= 1
	// break continue 一样的 但是break只能用在 loop,switch,select   continue只能在loop

	// for {} 相当于 while 
	 for {
		 i++
		 time.Sleep(time.Second)

		 if(i==8){
			 break
		 }
		fmt.Printf("i == %d \n",i);
	 }
	 fmt.Printf("test \n");

	 // 改变代码执行顺序，不能跨函数
	 goto End
	 fmt.Printf("01 \n");

	 End:
	 fmt.Printf("End \n");

}