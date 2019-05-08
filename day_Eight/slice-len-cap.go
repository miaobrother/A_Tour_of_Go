package main

import "fmt"

func main() {
	abc := []int{2, 3, 5, 7, 11, 13, 17, 19}
	printSlice(abc)

	abc = abc[:2]
	printSlice(abc)

	abc = abc[:4]
	printSlice(abc)

	abc = abc[2:]
	printSlice(abc)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
