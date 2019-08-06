package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := YourFloat(-math.Sqrt2)
	v := Ver{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = &v

	fmt.Println(a.Abs())
}

type YourFloat float64

func (f YourFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Ver struct {
	X, Y float64
}

func (v *Ver) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}