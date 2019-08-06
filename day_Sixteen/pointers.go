package main

import "fmt"

func main() {
	i, j := 41, 2701
	p := &i         // 指向i
	fmt.Println(*p) // 通过指针读取 i的值
	*p = 21         // 通过指针设置i的值
	fmt.Println(i)  // 查看i的值

	p = &j         // 指向i
	*p = *p / 37   // 指向指针对j进行除法运算
	fmt.Println(j) // 查看j的值
}
