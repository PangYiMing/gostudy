package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	env := "t1"                                           // "staging" , prod : ""
	baseUrl := "http://fe-domain-admin.prod.learnta.work" // http://fe-domain-admin.learnta.work/
	WriteWithIoutil("aa.js", "if (process.env.NODE_ENV !== 'development') {\n  process.env.VUE_APP_LERNTA_BACKEND = 'https://"+env+".learnta.cn';\n  process.env.BASEURL = '"+baseUrl+"';\n  process.env.VUE_APP_BASEURL = '"+baseUrl+"';\n}\n")
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("写入文件成功:", content)
	}
}
