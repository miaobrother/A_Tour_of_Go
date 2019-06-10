package main

import "fmt"

type VertexT struct {
	Z int
	N int
}

func main() {
	V := VertexT{1, 2}
	V.Z = 4
	fmt.Printf("The New VertexT is %v\n", V)
}
