package main
import "fmt"

func main ()  {
	var t complex128 //声明
	t = 2.1 + 3.14i

	fmt.Println("t =", t)

	t1 := 3.1 + 3.14i
// %T 复数

	fmt.Println("t1 =", real(t1), imag(t1))

}