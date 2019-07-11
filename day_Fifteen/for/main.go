package main

import "fmt"

func testFor() {
	var i int
	for i := 0; i <= 10; i++ {
		fmt.Printf("i = %d\n", i)
	}
	fmt.Println("The i is ", i)
}
func testDouble() {
	for no, i := 10, 1; no <= 19 && i <= 10; i, no = i+1, no+1 {
		fmt.Printf("%d * %d= %d\n", no, i, no*i)
	}
}
func main() {
	//testFor()
	testDouble()
}
