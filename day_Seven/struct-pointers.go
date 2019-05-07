package main

import "fmt"

type Nexus struct {
	Host string
	Port int
}

func main() {
	n := Nexus{"192.168.5.25", 9001}
	p := &n
	p.Host = "192.168.5.21"
	fmt.Println(n)
}
