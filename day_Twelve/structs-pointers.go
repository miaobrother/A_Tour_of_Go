package main

import "fmt"

type VertexN struct {
	A int
	B int
}

func main() {
	v := VertexN{1, 2}
	p := &v
	p.A = 1e9
	fmt.Printf("The p is : %v\n", *p)
}
