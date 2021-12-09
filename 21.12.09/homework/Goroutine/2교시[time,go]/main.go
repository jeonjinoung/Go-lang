package main

import (
	"fmt"
	"time"
)

//전역변수선언
var (
	//변수명 dataSendChannel = chan 배열 type int
	dataSendChannel = make(chan int)
)
//first 함수명선언 요소로 c변수선언 그값으로 chan int가져옴
func first(c chan<- int) {
	//for문 index가 1이라고 초기값을 잡았을때 인덱스가 10보다 작을때 하나씩 늘려준다.
	for i := 1; i <= 10; i++ {
		//그리고 second함수 실행 그안에 chan int넣어줌
		go second(dataSendChannel)
		//c라는 변수는 인덱스가 늘어날때마다 빼준다.
		c <- i
	}
}
//second 함수명선언 요소로 c <-chan int값
func second(c <-chan int) {
	//num변수선언 c보다 작을시 하나씩 빼준다. 
	num := <-c
	//time패키지에 Sleep 딜레이시켜준다 그리고 그값을 세컨드와 곱해준다.
	time.Sleep(time.Duration(10-num) * time.Second)
	//그리고 프린트해준다 num을
	fmt.Println("[NUM] : ", num)
}

func main() {
	// first함수는 chan int값을 넣어주고
	go first(dataSendChannel)
	//time패키지 딜레이시켜준다. 10초동안
	time.Sleep(time.Second * 100)
}

