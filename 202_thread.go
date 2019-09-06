package main
import (
    "runtime"
    "fmt"
)
// https://blog.csdn.net/coolwriter/article/details/80852639
func main(){
	// 加go用线程
    go sayHello()
    go sayWorld()
    var str string
    fmt.Scan(&str)
}
func sayHello(){
    for i := 0; i < 10; i++{
		fmt.Print("hello ")
		// 释放cpu
        runtime.Gosched()
    }
}
func sayWorld(){
    for i := 0; i < 10; i++ {
        fmt.Println("world")
        runtime.Gosched()
    }
}