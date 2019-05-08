package main

import "fmt"

func main() {
	var sli []int
	fmt.Println(sli, len(sli), cap(sli))
	if sli == nil {
		fmt.Println("Nil")
	}
}
