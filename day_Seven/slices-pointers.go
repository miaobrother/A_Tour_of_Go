package main

import "fmt"

func main() {
	names := [4]string{
		"Jone",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Printf("The name is: %v\n", names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println("The a :", a)
	fmt.Println("The b :", b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println("nes names is:", names)
}
