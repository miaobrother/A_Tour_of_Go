package main

import "fmt"

type JobStep struct {
	X int
	Y int
}

func main() {
	j := JobStep{}
	j.X = 15
	fmt.Println(j.X, j.Y)
}
