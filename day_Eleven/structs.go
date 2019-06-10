package main

import "fmt"

type Vertex struct {
	X string
	Y int
}

func main() {
	fmt.Printf("The Vertex is %v\n", Vertex{"abc", 1})
}
