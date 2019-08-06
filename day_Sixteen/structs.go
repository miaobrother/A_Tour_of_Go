package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式的赋值
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} //创建一个*Vertex类型的结构体指针

)

func main() {
	fmt.Println(v1, v2, v3, p)
}
