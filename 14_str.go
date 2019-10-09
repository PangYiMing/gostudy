package main
import "fmt"

func main ()  {
	var ch byte
	ch = 97
	
	fmt.Println("ch =", ch);
	fmt.Printf("%c ,  %d\n", ch, ch);
	ch = 97
	fmt.Printf("%c ,  %d\n", ch, ch);

	fmt.Printf("%d ,  %d\n", 'A', 'a');
	fmt.Printf("%d \n", 'A'+ 32);

	str := "myss"
	fmt.Println("str =", str, len(str))


	str2 := `sss
	\n
	`
	fmt.Println("str2 =", str2)

	str3 := str2+ str
	fmt.Println("str3 =", str3)


}