package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Product _
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}
type Auth struct {
	Username string `json:"username"`
	Pwd      string `json:"password"`
}

func getObjFromString(data string) *Product {
	p := &Product{}
	json.Unmarshal([]byte(data), p)
	return p
}
func (p *Product) setName(name string) *Product {
	p.Name = name
	return p
}

// 以下 json: cannot unmarshal string into Go value of type main.Product
// p1 := &Product{}
// data12, _ := json.Marshal(data)
// err = json.Unmarshal(data12, p1)
// fmt.Println(err)
// fmt.Println(p1.Name)

func getKeyValueFromStruct(user Auth, key string) string {
	// 从struct中取值
	immutable := reflect.ValueOf(user)
	val := immutable.FieldByName(key).String()
	return val
}

func main() {
	// 从json string中取值
	p := getObjFromString(`{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`)
	p.setName("sss")
	fmt.Println("print p:", p.Name)

	UserAdmin := Auth{
		Username: "admin",
		Pwd:      "123456",
	}

	// fmt.Printf("Username=%v\n", getKeyValueFromStruct(UserAdmin, "Username")) // prints 1

	// 打印struct
	// data1, _ := json.Marshal(UserAdmin)
	fmt.Println(UserAdmin.Username)
	// fmt.Println(string(data1))

}
