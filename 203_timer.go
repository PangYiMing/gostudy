package main

import (
    "fmt"
	"time"
    "runtime"
)

func main() {

	go sleep(1);
	go sleep(2);
	go print();
	var str string
    fmt.Scan(&str)


    // fmt.Printf("3\n")
    // //如果只是想单纯的等待的话，可以使用 time.Sleep 来实现
    // timer2 := time.NewTimer(time.Second * 2)
    // <-timer2.C
    // fmt.Println("2s后")

    // time.Sleep(time.Second * 2)
    // fmt.Println("再一次2s后")

    // <-time.After(time.Second * 2) //time.After函数的返回值是chan Time
    // fmt.Println("再再一次2s后")

    // timer3 := time.NewTimer(time.Second)
    // go func() {
    //     <-timer3.C
    //     fmt.Println("Timer 3 expired")
    // }()

    // stop := timer3.Stop() //停止定时器
    // ////阻止timer事件发生，当该函数执行后，timer计时器停止，相应的事件不再执行
    // if stop {
    //     fmt.Println("Timer 3 stopped")
    // }

    // fmt.Println("before")
    // timer4 := time.NewTimer(time.Second * 5) //原来设置5s
    // timer4.Reset(time.Second * 1)            //重新设置时间,即修改NewTimer的时间
    // <-timer4.C
    // fmt.Println("after")
}

func sleep(a int)  {
	for index := 0; index < 10; index++ {
        runtime.Gosched()
		t1 := <-time.NewTimer(time.Second * 2).C
		fmt.Printf("%d  %v\n",a, t1)
	}
}
func print()  {
	fmt.Printf("3\n")
	runtime.Gosched()
}