package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	dataSendChannel = make(chan int)
	rwMutex         = new(sync.RWMutex)
)

func main() {
	go first(dataSendChannel)
	time.Sleep(time.Second * 100)
}

func first(c chan<- int) {
	go second(dataSendChannel)
	go third(dataSendChannel)
	for i := 1; i <= 10; i++ {
		c <- i
	}
}

func second(c <-chan int) {
	rwMutex.Lock()
	fmt.Println("[NUM] : ", <-c)
	time.Sleep(time.Second * 1000)
	rwMutex.Unlock()
}

func third(c <-chan int) {
	for i := 1; i < 10; i++ {
		fmt.Println("[NUM] : ", <-c)
	}
}