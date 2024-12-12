package main

import (
	"time"
)

func main() {

	//1. Unbuffered Channel
	unbufferedChan := make(chan string)
	go func() {
		unbufferedChan <- "test"
		time.Sleep(1 * time.Second)
		unbufferedChan <- "test 2"
		close(unbufferedChan)
	}()
	println("Start")
	for v := range unbufferedChan {
		println(v)
	}
	println("End")

	//2. Buffered channel example
	bufferedChannel := make(chan int, 3)
	go func() {
		println("Start")
		bufferedChannel <- 1
		bufferedChannel <- 2
		bufferedChannel <- 3
		bufferedChannel <- 4
		close(bufferedChannel)
		//bufferedChannel <- 3 // will panic sending on close channel
	}()
	println("Mid")
	for v1 := range bufferedChannel {
		println(v1)
	}

	//3. Select using channel

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "Data from chan1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chan2 <- "Data from chan2"

	}()

	println("Start of select")
	select {
	case msg1 := <-chan1:
		println(msg1)
	case msg2 := <-chan2:
		println(msg2)
	case <-time.After(10 * time.Second):
		println("Timeout Occures")
	}
	println("End of select")

}
