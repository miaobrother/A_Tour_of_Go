package main

import "fmt"

func main() {
	var ss []int
	printSlicess(ss)
	ss = append(ss, 0)
	printSlicess(ss)
	ss = append(ss, 1)
	printSlicess(ss)

	ss = append(ss, 2, 3, 4, 5, 6)
	printSlicess(ss)

}

func printSlicess(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
