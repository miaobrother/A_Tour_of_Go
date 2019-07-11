package main

import "fmt"

func evenOdd() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("num %v is even\n", num)
	} else {
		fmt.Printf("num %v is odd\n", num)
	}
}

func testOdd() {
	num := 19
	if num >= 5 && num <= 10 {
		fmt.Printf("num is %v\n", num)
	} else if num >= 10 && num <= 20 {
		fmt.Printf("num %v is > 10 且 < 20\n", num)
	} else {
		fmt.Printf("%v\n", num)
	}
}

func test1Odd() {
	if num := getNum(); num %2 == 0 {
		fmt.Printf("num is %d\n", num)
	} else {
		fmt.Println(num)
	}
}

func getNum() int {
	return 10
}

func main() {
	evenOdd()
	testOdd()
	test1Odd()
}
