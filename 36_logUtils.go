package main

import (
	"fmt"
	"log"
	"path/filepath"
	_ "reflect"
	"runtime"
	"strings"
)

// https://studygolang.com/articles/9158
// https://www.jianshu.com/p/7aa54489bb4b
type MyStruct struct {
}

func (m *MyStruct) foo(p string) {
	HOSTCONFIGKEY := "sssss1"
	val := "sssss2"
	token := "tokensssss"
	ENTRY("")
	ENTRY("Param p=%s", p)
	DEBUG("Test %s %s", "Hello", "World")
	DEBUG("token: %v,redis  key = "+HOSTCONFIGKEY+" ,value = "+val, token)
}

func DEBUG(formating string, args ...interface{}) {
	LOG("DEBUG", formating, args...)
}

func ENTRY(formating string, args ...interface{}) {
	LOG("ENTRY", formating, args...)
}

func LOG(level string, formating string, args ...interface{}) {
	filename, line, funcname := "???", 0, "???"
	_, myfilePath, _, _ := runtime.Caller(1)
	pc, filename, line, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo

		filename = filepath.Base(filename) // /full/path/basename.go => basename.go
	}

	log.Printf("%s:%d:%s: %s: %s\n", myfilePath, line, funcname, level, fmt.Sprintf(formating, args...))
}

func main() {
	ss := MyStruct{}
	ss.foo("helloworld")

}
