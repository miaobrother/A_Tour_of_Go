package main

import "fmt"

func main() {
	var names []string

	PrintSlice(names)

	names = append(names, "One")
	PrintSlice(names)

	names = append(names, "Two")
	PrintSlice(names)

	names = append(names, "Three Four Five")
	PrintSlice(names)
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for i := range names {
		fmt.Println(names[i])
	}
}

func PrintSlice(s []string) {
	fmt.Printf("The len=%d cap=%d %v\n", len(s), cap(s), s)
}
