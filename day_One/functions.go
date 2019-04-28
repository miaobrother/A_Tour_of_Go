package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	Results := add(10, 20)
	fmt.Println("The res is:", Results)
}
