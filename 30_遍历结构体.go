package main

import (
	"fmt"
	"reflect"
)

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

// HostConfig 域名列表
type HostConfig struct {
	T1      string `json:"t1"`
	T2      string `json:"t2"`
	T3      string `json:"t3"`
	Sit     string `json:"sit"`
	Staging string `json:"staging"`
	Rc      string `json:"rc"`
	Beta    string `json:"beta"`
	Pro     string `json:"pro"`
}

func main() {
	hConfig := EnvObj{
		T1:      0,
		T2:      3,
		T3:      3,
		Sit:     3,
		Staging: 3,
		Rc:      3,
		Beta:    3,
		Pro:     2,
	}
	myHostConfig := HostConfig{
		T1:      "{1}",
		T2:      "{2}",
		T3:      "{3}",
		Sit:     "{4}",
		Staging: "{5}",
		Rc:      "{6}",
		Beta:    "{7}",
		Pro:     "{8}",
	}
	filter(&myHostConfig, &hConfig)
	fmt.Printf("%v\n", myHostConfig)
}

func filter(myHostConfig *HostConfig, judgeConfig *EnvObj) {
	value := reflect.ValueOf(myHostConfig)
	key := reflect.TypeOf(myHostConfig)
	value2 := reflect.ValueOf(judgeConfig)
	key2 := reflect.TypeOf(judgeConfig)
	if key.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		key = key.Elem()
		value = value.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}
	if key2.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		key2 = key2.Elem()
		value2 = value2.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}
	v := "{}"
	for i := 0; i < value.NumField(); i++ {
		val := value.Field(i)
		k := key.Field(i)
		val2 := value2.Field(i)
		// k2 := key2.Field(i)
		convertValue := val2.Interface().(int)
		if convertValue == 0 {
			fmt.Printf("Field index %d: key: %v, value: %v \n", i, k.Name, val)
			dataType := reflect.TypeOf(v)
			structType := val.Type()
			fmt.Printf("type  dataType %v: structType: %v \n", dataType, structType)
			if structType == dataType {
				val.Set(reflect.ValueOf(v))
				fmt.Printf("second Field index %d: key: %v, value: %v \n", i, k.Name, val)
			} else {
				if dataType.ConvertibleTo(structType) {
					// 转换类型
					val.Set(reflect.ValueOf(v).Convert(structType))
					fmt.Printf("thirst Field index %d: key: %v, value: %v \n", i, k, val)
				}
			}
		}

	}
}
