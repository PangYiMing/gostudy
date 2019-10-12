package main

import (
	"fmt"
	"strings"
)

func main() {
	var ch byte
	ch = 97

	fmt.Println("ch =", ch)
	fmt.Printf("%c ,  %d\n", ch, ch)
	ch = 97
	fmt.Printf("%c ,  %d\n", ch, ch)

	fmt.Printf("%d ,  %d\n", 'A', 'a')
	fmt.Printf("%d \n", 'A'+32)

	str := "myss"
	fmt.Println("str =", str, len(str))

	str2 := `sss
	\n
	`
	fmt.Println("str2 =", str2)

	str3 := str2 + str
	fmt.Println("str3 =", str3)
	if 2 == 2 && strings.EqualFold("GO", "go") {
		fmt.Println("true")
	}
	// Go语言---strings包(字符串操作)
	// https://blog.csdn.net/li_101357/article/details/80241224
	pre := "\""
	dis := "\""
	var hostName string
	testStr := "\"test\""
	host := "test"
	hostName = pre + host + dis

	fmt.Println(strings.Count(testStr, hostName))
}
