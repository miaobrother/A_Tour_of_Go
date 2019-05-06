package main

import (
	"fmt"
	"math"
)

func square(x float64) string {
	if x < 0 {
		return square(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(square(2), square(-4))

}
