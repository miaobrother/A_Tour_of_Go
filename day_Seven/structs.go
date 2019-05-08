package main

import "fmt"

type Info struct {
	Age  int
	Name string
}

func main() {
	fmt.Println(Info{28, "Zhangsan"})
}
