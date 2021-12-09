package main

import (
	"fmt"
	"time"
)

var (
	dataSendChannel = make(chan int)
)

func first(c chan<- int) {
	for i := 1; i <= 10; i++ {
		go second(dataSendChannel)
		c <- i
	}
}

func second(c <-chan int) {
	fmt.Println("[NUM] : ", <-c)
}

func main() {
	go first(dataSendChannel)
	time.Sleep(time.Second * 10)
}
