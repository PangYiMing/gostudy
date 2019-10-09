package main

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
	"reflect"
)
var c *redis.ClusterClient
func main() {
	c = redis.NewClusterClient(&redis.ClusterOptions{
        Addrs:    []string{"127.0.0.1:9001"}, //set redis cluster url
    })
	fmt.Println("type:", reflect.TypeOf(c))

	defer c.Close()

	result := c.Do("SET", "mykey", "superWang", "EX", "5")
	fmt.Println("redis set result:", result)

	username := c.Do("GET", "mykey")
		fmt.Printf("Get mykey: %v \n", username)

	time.Sleep(2 * time.Second)

	// pong, err := c.Ping().Result()
    // fmt.Println(pong, err)
    // fmt.Println("pool state init state:", c.PoolStats())
    for i := 0; i < 1000; i++ {
		// Sprintf 格式化输出
        k := fmt.Sprintf("key:%d", i)
        v := k
        val, err := c.Set(k, v, 60*time.Second).Result()
        if err != nil {
            panic(err)
        }
 
        val, err = c.Get(k).Result()
        if err != nil {
            panic(err)
        }
        fmt.Println("key:", val)
    }
    fmt.Println("pool state final state:", c.PoolStats()) //获取客户端连接池相关信息
 
}