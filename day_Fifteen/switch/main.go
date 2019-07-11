package main

import "fmt"

func testSwitch() {
	finger := 4
	switch finger {
	case 1, 4:
		{
			fmt.Println("The finger is Thumb")
		}
		fallthrough
	case 2:
		{
			fmt.Println("The finger is  Index")
		}
	case 3:
		{
			fmt.Println("The finger is Middle")
		}
	default:
		{
			fmt.Println("Nnknow data")
		}

	}

}

func test1Switch()  {

}

func main() {
	testSwitch()

}
