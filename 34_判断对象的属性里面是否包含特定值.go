package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// var a int = 10

	adminEnv := EnvObj{
		T1:      3,
		T2:      3,
		T3:      3,
		Sit:     3,
		Staging: 3,
		Rc:      3,
		Beta:    3,
		Pro:     3,
	}
	testStr := "T1"
	bl := judgeEnvPermision(&adminEnv, testStr)
	fmt.Println("last last ", bl)
}

// 0 没有权限， 1只有读权限， 2 只有写权限 ，3 读写权限
type EnvObj struct {
	T1      int `json:"t1"`
	T2      int `json:"t2"`
	T3      int `json:"t3"`
	Sit     int `json:"sit"`
	Staging int `json:"staging"`
	Rc      int `json:"rc"`
	Beta    int `json:"beta"`
	Pro     int `json:"pro"`
}

func judgeEnvPermision(myHostConfig *EnvObj, env string) bool {
	value := reflect.ValueOf(myHostConfig)
	key := reflect.TypeOf(myHostConfig)
	fmt.Println("string ", env)

	if key.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		key = key.Elem()
		value = value.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}

	for i := 0; i < value.NumField(); i++ {
		val := value.Field(i)
		k := key.Field(i)
		fmt.Printf("Field index %d: key: %v, value: %v \n", i, k.Name, val)
		currentIndex := val.Interface().(int)
		if currentIndex >= 2 && strings.EqualFold(k.Name, env) {
			fmt.Println("ssss true")
			return true
		}
	}
	return false
}
