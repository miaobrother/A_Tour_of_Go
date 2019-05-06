package main

import (
	"fmt"
	"math"
)

func Pow2(x, y, limit float64) float64 {
	if v := math.Pow(x, y); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit
	// can't use v here
}

func main() {
	fmt.Println(
		Pow2(2, 3, 9),
		Pow2(10, 3, 1001),
	)

}
