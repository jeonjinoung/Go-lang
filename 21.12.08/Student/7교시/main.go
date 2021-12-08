package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)//요소가 5개인 링 생성

	n := r.Len()//길이는 5

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r=r.Next()
	}
	////////////////////////////////////////
	//역순
	for k := 0; k < n; k++ {
		fmt.Printf("%c",r.Value)
		r=r.Prev()
	}
	///////////////////////////////////////
	//정순
	// for k := 0; k < n; k++ {
	// 	fmt.Printf("%c",r.Value)
	// 	r=r.Next()
	// }
}

/*
링(ring) -> 원형으로 연결되어있다. -> 반복된다. 환영큐 서클라큐
마지막 요소가 처음요소까지 연결되어있다.

다른언어에서는 native로 만들어줘야한다.
container안에 있다.

링은 언제쓰냐? 실행취소할때?? 리플레이기능 구현시
*/