package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	get()
	// sleep(1)
	// post();
}

func get() {
	response, err := http.Get("http://127.0.0.1:5000/__render/updateHost")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.StatusCode)
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))
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
	_, newerr := ioutil.ReadAll(response.Body)
	if newerr != nil {
		return
	}
	// fmt.Println(string(body))
	fmt.Printf("01")
}

func sleep(a int) {
	for index := 0; index < 30; index++ {
		// runtime.Gosched()
		t1 := <-time.NewTimer(time.Second * 2).C
		fmt.Printf("%d  %v\n", a, t1)
		for i := 1; i <= 300; i++ {
			post()
		}
	}
}
