package main

import (
	"fmt"
	"time"
)

func testHuiWen() {
	var str = "上海自来水来自海上"
	var r = []rune(str)
	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[i] = tmp
	}
	str2 := string(r)
	if str2 == str {
		fmt.Println("饰回纹")
	} else {
		fmt.Println("不是会问")
	}

}

func testTimeStamp(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	year := timeObj.Year()
	fmt.Printf("The year: %v\n", year)

}

func testTime() {
	now := time.Now()
	fmt.Println("now: ", now)
	year := now.Year()
	fmt.Println(year)
	month := now.Month()
	fmt.Println(month)
	fmt.Println(now.Location())
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

}

func processTask() {
	fmt.Println("wo ciao ")
}

func testTimer() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Printf("current time is %v\n", i)
		processTask()
	}
}

func main() {
	//testHuiWen()
	testTime()
	a := time.Now().Unix()
	testTimeStamp(a)
	testTimer()
}
