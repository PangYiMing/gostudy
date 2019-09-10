package main
import "fmt"

func main ()  {
	var a [10]int
	var b [7]int
	fmt.Printf("%d ,%d \n",len(a),len(b))

	a[0] = 1
	fmt.Printf("%d \n",a[0])


	//声明定义同时赋值，叫初始化 
	// 全部初始化
	var c [5]int = [5]int{1,2,3,4,4}
	fmt.Println("c = ",c)
	// 部分初始化，没有初始化的数值，自动赋值0
	d := [5]int{1,2,3}
	fmt.Println("d = ",d)

	// 部分初始化，指定某个元素初始化
	e := [5]int{1:2,3:4}
	fmt.Println("e = ",e)

	// 二维数组
	// 有多少个【】就有多少维，对应多少循环
	var f [3][4]int
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			f[i][j] = k
			fmt.Printf("f[%d][%d] = %d \n",i,j,f[i][j])
		}
	}
	fmt.Println("f = ",f)

	g := [3][4]int{
		{1 ,2 ,3 ,4},
		{5 ,6 ,3 ,4},
		{1 ,0 ,3 ,4},
	}
	fmt.Println("g = ",g)
	j := [3][4]int{
		{1 ,2 ,3 ,},
		{5 ,6 ,3 ,4},
		{1 ,0 },
	}
	fmt.Println("j = ",j)

	h := [3][4]int{
		1:{5 ,6 ,3 ,4},
		{1 ,0 },
	}
	fmt.Println("h = ",h)


	// 数组比较,只支持 == 或 !=,比较是不是每一个元素都一样，2个数组比较数组类型要一样
	fmt.Println("c = ",c, ", d = ",d, ", e = ",e)
	// 同一类型可以赋值
	c = e
	fmt.Println("c == e ",c==e)


	modify(c) // 数组传递过去，拷贝后传出去， 值传递
	fmt.Println("main: c = ",c)
	modify2(&c) // 地址传递
	fmt.Println("main: c = ",c)


}
// 形参是实际参数的复制品
func modify(a [5]int)  {
	a[0] = 11
	fmt.Println("main: c = ",a) 
}

// 指向地址 *a是数组指针
func modify2(a *[5]int)  {
	(*a)[0] = 11
	fmt.Println("main: c = ",a)
}