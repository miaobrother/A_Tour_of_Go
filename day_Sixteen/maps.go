package main

import "fmt"

type Vertext struct {
	Lat, Long float64
}

var m map[string]Vertext

func main() {
	m = make(map[string]Vertext)
	m["Bell Labs"] = Vertext{
		40.68, -74.12111,
	}
	fmt.Println(m)
}
