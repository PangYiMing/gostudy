package main

import (
	"encoding/json"
	"fmt"
)

// 设置域名每一项
type HostItem struct {
	HtmlName  string   `json:"htmlName"`
	Host01    string   `json:"host01"`
	Host02    string   `json:"host02"`
	HostDIY   []string `json:"hostDIY"`
	WWWDelete bool     `json:"wwwDelete"`
}

func main() {
	// origin1 := `[{"htmlName":"learnta1","host01":"learnta1.cn","host02":"","hostDIY":null,"wwwDelete":false}]`
	origin2 := `[{"htmlName":"learnta2","host01":"learnta2.cn","host02":"learnta2.cn"}]`
	// var hosts []HostItem

	// item := HostItem{
	// 	Host01:   "ss",
	// 	HtmlName: "ss",
	// }
	// data12, _ := json.Marshal(item)
	// fmt.Println("strAndArr Marshal:", string(data12))
	// hosts = append(hosts, item) // 给切片的末尾增加一个成员,cap 不变
	// arrayData, _ := json.Marshal(hosts)
	// fmt.Println("strAndArr Marshal:", string(arrayData))

	// var hosts2 []HostItem
	// json.Unmarshal([]byte(origin1), &hosts2)
	// fmt.Println("strAndArr hosts2:", hosts2, len(hosts2))

	var hosts3 []HostItem
	json.Unmarshal([]byte(origin2), &hosts3)
	fmt.Println("strAndArr hosts3:", hosts3, len(hosts3))

	strAndArr(hosts3)

}

func strAndArr(hosts []HostItem) {

	fmt.Println("strAndArr fun:", hosts)
}
