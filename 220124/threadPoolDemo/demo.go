package main

import (
	"fmt"
	"math/rand"
	"time"
)

func taskExecutor(c, ret chan int) {
	for {
		taskID, ok := <-c
		if !ok {
			ret <- 1 // Indicating this thread has finished
			return
		}
		timeCost := time.Duration(30000.0*rand.Float64()) * time.Second / 10000
		fmt.Println("Executing task #", taskID, ", cost", timeCost.Milliseconds(), "ms")
		time.Sleep(timeCost)
		fmt.Println("Finished task #", taskID)
	}
}

func main() {
	const threadCount = 3
	taskQueue := make(chan int)
	retChannel := make(chan int)
	for i := 0; i < threadCount; i++ {
		go taskExecutor(taskQueue, retChannel)
	}

	for i := 0; i < 10; i++ {
		taskQueue <- i
	}
	close(taskQueue)

	// Wait for all threads to finish
	for i := 0; i < threadCount; i++ {
		_ = <-retChannel
	}
}
