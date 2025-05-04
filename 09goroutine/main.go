package main

import (
	"fmt"
	"time"
)

func doActivity(s string) {
	for i := 0; i < 11; i++{
		fmt.Println(s, " ", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go doActivity("drinking")
	go doActivity("eating")

	// add sleep so main doesn't finish before goroutines can run
	time.Sleep(6 * time.Second)
}