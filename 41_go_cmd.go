package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		var nativeEnv string = "native"
		if strings.EqualFold(os.Args[1], nativeEnv) {
			fmt.Println(os.Args[1])
		}
	}
}
