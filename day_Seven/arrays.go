package main

import "fmt"

func main() {
	var i [10]int
	i[0] = 1
	i[2] = 2
	fmt.Println(i)

	primes := [8]int{2, 3, 5, 7, 11, 13, 17, 19}
	fmt.Println(primes)
}
