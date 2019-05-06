package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("The Number is:",rand.Intn(10))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
