// 


package main
import "math/rand"
import "time"
import (
	"fmt"
)

func main(){
	array :=  []int{1,2,3,0,0,0}
	s := array[0:3:5] // 左闭右开 [low:high:max] low 下标起点 high 终点 max 最大的容量
	fmt.Println("s =", s)
	fmt.Println("len(s) =", len(s))
	fmt.Println("cap(s) =", cap(s))

	s = array[1:4:4] // 左闭右开  nvalid slice index: 4 > 3  max-low = cap >high-low
	fmt.Println("s =", s)
	fmt.Println("len(s) =", len(s))
	fmt.Println("cap(s) =", cap(s))

	// 数组的长度是固定的常量，不能修改
	// a := [5]int{}


	 // 切片 的 [] 为空，或者为...切片的长度可以不固定
	 u :=  []int{}
	 fmt.Printf("1: len = %d, cap = %d \n", len(u), cap(u)) 
	 u = append(array, 11) // 给切片的末尾增加一个成员
	 fmt.Println("s =", u)
	 fmt.Printf("1: len = %d, cap = %d \n", len(u), cap(u))  

	 // 自动推导类型
	 s1 :=  []int{1, 2, 3, 5}
	 fmt.Println("s1=",s1)

	// 借助 make 函数，格式make(slice,len,cap)
	s2 := make([]int, 5, 9)
	s3 := make([]int, 5)
	// cap...????
	fmt.Println("s2=",s2, len(s2), cap(s2))
	fmt.Println("s3=",s3, len(s3), cap(s3))
	s3 = append(s3, 11) // 给切片的末尾增加一个成员,cap 变10
	fmt.Println("s3=",s3, len(s3), cap(s3))
	s2 = append(s2, 11) // 给切片的末尾增加一个成员,cap 不变
	fmt.Println("s2=",s2, len(s2), cap(s2))

	s3 = array[:2] // 从0开始取3个元素，容量与array一致
	fmt.Println("s3=",s3, len(s3), cap(s3))
	s3 = array[2:] // 从3开始取到结尾，容量是长度
	fmt.Println("s3=",s3, len(s3), cap(s3))

	s3 = array[:] // 全复制，容量，长度一致 
	fmt.Println("s3=",s3, len(s3), cap(s3))
 

	//切片和底层数组的关系 切片的指针指向底层数组
	newA := []int{0,1,2,3,4,5,6,7,8,9}

	newB := append(newA,10) // 新的数组，与老得无关
	newB[8] = 99
	fmt.Println("newA=",newA, len(newA), cap(newA))// newA= [0 1 2 3 4 5 6 7 8 9] 10 10
	fmt.Println("newB=",newB, len(newB), cap(newB))// newB= [0 1 2 3 4 5 6 7 99 9 10] 11 20
	newB = newA[2:5] // 取3个
	fmt.Println("newB=",newB, len(newB), cap(newB))// newB= [2 3 4] 3 8
	newC := newB[2:7] // 这样取内容的时候，newB 的容量必须够 panic: runtime error: slice bounds out of range
	fmt.Println("newC=",newC, len(newC), cap(newC)) // newC= [4 5 6 7 8] 5 6
	newC[2] = 777
	newB[1] = 666
	fmt.Println("newA=",newA, len(newA), cap(newA))
	fmt.Println("newB=",newB, len(newB), cap(newB))
	fmt.Println("newC=",newC, len(newC), cap(newC))

	// append 超过以后以两倍的大小扩容


	copy(newC,newB) // newB -》 newC
	fmt.Println("newB=",newB, len(newB), cap(newB))
	fmt.Println("newC=",newC, len(newC), cap(newC))



	// 切片本身就是指针，不用拷贝
	newN := 10
	newS := make([]int, newN)
	InitData(newS)
	fmt.Println("newS=",newS, len(newS), cap(newS))

	BubbleSort(newS)
	fmt.Println("newS=",newS, len(newS), cap(newS))

}
func InitData(s []int)  {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
}
func BubbleSort(s []int)  {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j],s[j+1] = s[j+1],s[j] 
			}
		}
	}
}