package main

import "fmt"

func main() {
	s3 := []int{1, 2, 3, 4, 5, 6}
	s3 = s3[1:]
	fmt.Println(s3)
	s3 = s3[:2]
	fmt.Println(s3)
}
