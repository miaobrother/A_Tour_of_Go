package main

import (
	"fmt"
	"math"
)

func Pow(x, y, limit float64) float64 {
	if v := math.Pow(x, y); v < limit {
		return v
	}
	return limit
}

func main() {
	fmt.Println(
		Pow(3, 2, 10),
		Pow(20, 2, 402),
	)
}
