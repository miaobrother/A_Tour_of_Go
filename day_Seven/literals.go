package main

import "fmt"

type Vertex struct {
	Weight float64
	Height int
}

var (
	v1 = Vertex{125.1, 168}
	v2 = Vertex{Weight: 126.1}
	v3 = Vertex{}
	p  = &Vertex{127.3, 170} // 创建一个  *Vertex 类型的结构体（指针）
)

func main() {
	fmt.Println(v1, v2, v3, p)

}
