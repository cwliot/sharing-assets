package main

import (
	"fmt"
	"time"
)

func printer(c, ret chan int) {
	for {
		x, ok := <-c
		if !ok {
			break
		}
		fmt.Println(x)
	}
	fmt.Println("Channel closed")
	ret <- 999
	fmt.Println("Returning")
}

func main() {
	channel := make(chan int)
	ret := make(chan int)
	go printer(channel, ret)
	for i := 0; i < 5; i++ {
		channel <- i
		time.Sleep(1 * time.Second)
	}
	close(channel)
	time.Sleep(1 * time.Second)
	// x := <-ret
	// fmt.Println("printer told us:", x)
	time.Sleep(100 * time.Millisecond) // Allow time for the goroutine to finish
}
