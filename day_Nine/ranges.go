package main

import "fmt"

var powS = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range powS {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
