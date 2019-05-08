package main

import (
	"fmt"
)

func main() {
	q := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(q)

	r := [3]bool{true, false, true}
	fmt.Println(r)

	s := &[]struct {
		i int
		b bool
	}{
		{2, false}, {3, true},
	}
	fmt.Println(s)

}
