package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type IPAddr [4]byte

func (ip IPAddr) String() string  {
	
}

func (p Person) String() string  {
	return fmt.Sprintf("%v (%v years)",p.Name,p.Age)
}

func main() {
	a := Person{"Author",42}
	z := Person{"Zaphod",9000}
	fmt.Println(a,z)
}
