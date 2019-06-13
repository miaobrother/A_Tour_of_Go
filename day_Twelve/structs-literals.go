package main

import "fmt"

type Circle struct {
	O, P int
}

var (
	c1 = Circle{1, 2}
	c2 = Circle{P: 1}
	c3 = Circle{}
	p  = &Circle{11, 22}
)

func main() {
	fmt.Println(c1, c2, c3, p)
}
