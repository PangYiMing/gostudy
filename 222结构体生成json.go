package main

import (
	"encoding/json"
	"fmt"
)

type IT struct { //一定要注意这里的成员变量的名字首字母必须是大写
	Company  string
	Subjects []string
	Isok     bool
	Price    float64
}
type DATD struct { //一定要注意这里的成员变量的名字首字母必须是大写
	Username   string `json:"username"`
	Code       string `json:"code"`
	SystemCode string `json:"systemCode"`
}

// 首字母必须大写
func main() {
	s := IT{"zyg", []string{"go", "python", "java"}, true, 3.14} //json的初始化
	buf, err := json.Marshal(s)                                  //使用这个函数会返回两个值，通过源码得知
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf)) //注意这里生成的buf是一个byte切片，如果直接打印会是一串数字，这里使用string函数进行转化

	str := DATD{"18722530252", "1111", "domain"} //json的初始化
	fmt.Println("str = ", str)                   //注意这里生成的buf是一个byte切片，如果直接打印会是一串数字，这里使用string函数进行转化

	buf2, err2 := json.Marshal(str) //使用这个函数会返回两个值，通过源码得知
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	fmt.Println("buf2 = ", string(buf2)) //注意这里生成的buf是一个byte切片，如果直接打印会是一串数字，这里使用string函数进行转化

}
