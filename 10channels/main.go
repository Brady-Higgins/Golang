package main

import (
	"fmt"
	"time"
)

func randomFunc(intChannel chan int) {
	for i := 0; i < 20; i++ {
		// write to channel
		intChannel <- i
	}
	//close after done
	close(intChannel)
}

func evenOnly(intChannel chan int) {
	for {
		//read from channel
		n, isOpen := <-intChannel
		//channel doesn't auto stop once closed
		if !isOpen{
			break
		}
		if n%2 == 0 {
			fmt.Println(n)
		}
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	intChannel := make(chan int)
	go randomFunc(intChannel)
	evenOnly(intChannel)
}