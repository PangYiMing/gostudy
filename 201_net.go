package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func main(){
	// get();
	for i := 1; i <= 10; i++ {
		post();
    }
}

func get()  {
	response,err := http.Get("http://www.baidu.com")
	if(err!=nil){
		fmt.Println(err)
	}
	defer response.Body.Close()
	body,err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
func post()  {
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