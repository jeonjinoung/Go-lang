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

func first(c chan<- int) {
	for i := 1; i <= 10; i++ {
		go second(dataSendChannel)
		c <- i
	}
}

func second(c <-chan int) {
	rwMutex.Lock()
	num := <-c
	time.Sleep(time.Duration(10-num) * time.Second)
	fmt.Println("[NUM] : ", num)
	rwMutex.Unlock()
}

func main() {
	go first(dataSendChannel)
	time.Sleep(time.Second * 100)
}

