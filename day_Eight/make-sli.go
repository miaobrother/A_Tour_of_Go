package main

import "fmt"

func main() {
	a := make([]int, 10)
	printSli("a", a)

	b := make([]int, 0, 10)
	printSli("b", b)

	c := b[:2]
	printSli("c", c)
	printSli("b",b)

	d := c[2:]
	printSli("d",d)

}

func printSli(s string, x []int) {
	fmt.Printf("%s len= %d cap = %d %v\n", s, len(x), cap(x), x)
}
