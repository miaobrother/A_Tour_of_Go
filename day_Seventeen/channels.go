package main

import "fmt"

func sum(sli []int, c chan int) {
	sum := 0
	for _, v := range sli {
		sum += v
	}
	c <- sum
}

func main() {
	sli := []int{1, 2, 3, 4, 5, 6, 7}
	c := make(chan int)
	go sum(sli[:len(sli)/2], c)
	go sum(sli[len(sli)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
