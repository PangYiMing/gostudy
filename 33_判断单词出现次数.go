package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type H struct {
	A string `json:"a"`
	B string `json:"b"`
}

func main() {
	var str string = "hello hello oosss,hello.com,\"hello\","

	fmt.Println(strings.Count(str, "\"hello\""))

	myH := &H{}
	myH.A = "sf"
	data, _ := json.Marshal(myH)
	fmt.Println("/setHostConfig : set json data = ", string(data))

}
