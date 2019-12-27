package main

import (
	"fmt"
	"time"
)

func main() {

	timeUnix := time.Now().Format("2006-01-02&15:04:05")

	str := "backup" + fmt.Sprintf("%v&%v", timeUnix, time.Now().Unix())

	fmt.Printf("%s", str)
}
