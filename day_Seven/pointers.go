package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i // 指向 i
	fmt.Println("The j is:", &j)
	fmt.Printf("The p is: %v\n", *p) // 通过指针读取  i 的 值
	*p = 21                          // 通过指针设置 i 的值
	fmt.Println(i)
	p = &j
	fmt.Println(*p)
	*p = *p / 37
	fmt.Println(j)

}
