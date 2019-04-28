package main

import "fmt"

func swap(x, y string) (a, b string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
