package main

import "fmt"

func main() {

	// fmt.Println("4>3?", 4 > 3)

	fmt.Println("||", judge(1)|judge(2))

}
func judge(a int) bool {
	if a == 2 || a == 3 {
		fmt.Println("false", a)
		return false
	}
	fmt.Println("true", a)
	return true
}
