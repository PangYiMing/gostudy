package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	// get();
	for i := 1; i <= 10; i++ {
		// post()
	}
	myREQDATD := &REQDATD{"15000605075", "1111", "domain"} //json的初始化
	myEnvObj := &EnvObj{}
	var statusCode string = "-1"
	queryLDPNet(myREQDATD, myEnvObj, &statusCode)
	fmt.Println("statusCode", statusCode)
}

func get() {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
func post() {
	client := &http.Client{}
	requst, err := http.NewRequest("POST",
		"http://www.baidu.com",
		strings.NewReader("name=abc"))
	if err != nil {
		return
	}
	requst.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(requst)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}

type REQDATD struct { //一定要注意这里的成员变量的名字首字母必须是大写
	Username   string `json:"username"`
	Code       string `json:"code"`
	SystemCode string `json:"systemCode"`
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

func queryLDPNet(myREQDATD *REQDATD, myEnvObj *EnvObj, statusCode *string) {
	client := &http.Client{}

	buf2, err2 := json.Marshal(myREQDATD) //使用这个函数会返回两个值，通过源码得知
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}
	data := string(buf2)

	fmt.Println("data = ", data) //注意这里生成的buf是一个byte切片，如果直接打印会是一串数字，这里使用string函数进行转化

	url := "https://t1.learnta.cn/__api/auth/crm/user/login"
	requst, err := http.NewRequest("POST", url, strings.NewReader(data))

	if err != nil {
		return
	}
	requst.Header.Set("Content-Type", "application/json")
	response, err := client.Do(requst)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
	b := getObjFromString(string(body))
	if strings.EqualFold(b.Code, "0") {
		*statusCode = "0"
		myPrivileges := b.Data.USER.Privileges
		if setEnvObjPermissiom(myEnvObj, myPrivileges) {
			// TODO 其他
			fmt.Println("myEnvObj", myEnvObj)
		}

	}
}

type Privilege struct {
	Code string `json:"code"`
}
type PRIVILEGES struct {
	Privileges []Privilege `json:"privileges"`
}
type RESDATD struct {
	USER   PRIVILEGES `json:"user"`
	Name   string     `json:"name"`
	Mobile string     `json:"mobile"`
}
type LDPBody struct {
	Err        string  `json:"err"`
	Code       string  `json:"code"`
	Message    string  `json:"message"`
	Data       RESDATD `json:"data"`
	Successful bool    `json:"successful"`
}

func getObjFromString(data string) *LDPBody {
	p := &LDPBody{}
	json.Unmarshal([]byte(data), p)
	return p
}

// TODO log 选择时候带上token
func setEnvObjPermissiom(myEnvObj *EnvObj, myPrivileges []Privilege) bool {
	value := reflect.ValueOf(myEnvObj)
	key := reflect.TypeOf(myEnvObj)

	if key.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		key = key.Elem()
		value = value.Elem()
	} else {
		return false
	}
	fmt.Printf("myPrivileges i: %v\n", len(myPrivileges))

	for i := 0; i < len(myPrivileges); i++ {
		fmt.Printf("myPrivileges i: %v, Code: %v \n", i, myPrivileges[i].Code)
		code := myPrivileges[i].Code
		codeArray := strings.Split(code, ".")
		if len(codeArray) == 3 {

			for i := 0; i < value.NumField(); i++ {
				val := value.Field(i)
				k := key.Field(i)
				if strings.EqualFold(k.Name, codeArray[1]) {
					if strings.EqualFold("read", codeArray[2]) {
						var int1 int64 = 1
						permission := int1 + val.Int()
						// method 1:
						strInt64 := strconv.FormatInt(permission, 10)
						id16, _ := strconv.Atoi(strInt64)

						fmt.Print("permission: ", permission, ", val: ", val, ", id16: ", id16, "\n")
						val.Set(reflect.ValueOf(id16))
					} else {
						var int2 int64 = 2
						permission := int2 + val.Int()
						// method 1:
						strInt64 := strconv.FormatInt(permission, 10)
						id16, _ := strconv.Atoi(strInt64)
						fmt.Print("permission: ", permission, ", val: ", val, ", id16: ", id16, "\n")
						val.Set(reflect.ValueOf(id16))
					}
				}
			}

		}

	}

	return true
}
