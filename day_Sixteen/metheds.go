package main

import (
	"fmt"
	"math"
)

type Verte struct {
	X, Y float64
}

func (v Verte) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v *Verte) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Verte ,f float64)  {
	v.X =v.X *f
	v.Y = v.Y *f
}

func main() {
	v := Verte{3, 4}
	Scale(&v,10)
	fmt.Println(v.Abs())
	fmt.Println(Abs(&v))
}
