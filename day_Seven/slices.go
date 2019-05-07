package main

import "fmt"

func main() {
	Primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("The Primes type is %T and Primes value is:%v\n", Primes, Primes)
	newPrimes := Primes[1:len(Primes)-1]
	fmt.Printf("The newPrimes type is %T and newPrimes is %v\n", newPrimes, newPrimes)
}
