package main
import "fmt"
import "math/rand"
import "time"

func main ()  {
 
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		// rand.Int() 随机一个很大的数
		fmt.Println("rand = ", rand.Intn(100));
	}
	fmt.Println("4>3?", 4>3);


}