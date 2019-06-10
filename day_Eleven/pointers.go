package main

import "fmt"

func main() {

	i, j := 42, 2701
	p := &i                             // 指向 i
	fmt.Println(*p)                     // 通过指针读取 i 的值
	*p = 21                             // 通过指针设置 i 的值
	fmt.Printf("The new i is :%d\n", i) //  查看 i 的值

	p = &j
	*p = *p / 37
	fmt.Printf("The new j is :%d\n", j)
}
