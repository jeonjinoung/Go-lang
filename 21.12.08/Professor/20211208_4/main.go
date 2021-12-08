package main

import (
	"container/ring"
	"fmt"
)

type Product struct {
	Name  string
	Price int
}

func main() {

	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}
	// for k := 0; k < n; k++ {
	// 	fmt.Printf("%c", r.Value)
	// 	r = r.Next()
	// }
	// for k := 0; k < n; k++ {
	// 	fmt.Printf("%c", r.Value)
	// 	r = r.Prev()
	// }

	//map은 key와 value 쌍으로 이루어져 있다. 컨테이너 패키지가 아닌 기본 내장타입이다.
	//				키     value
	// m := make(map[string]string) //맵 생성

	// m["이창희"] = "서울시 강남구" //키와 값을 추가
	// m["바나나"] = "청송 교도소"  //키와 값을 추가
	// m["홍길동"] = "한라산"     //키와 값을 추가
	// m["김정은"] = "평양 어딘가"  //키와 값을 추가

	// m["바나나"] = "유치장으로 이동"
	// fmt.Printf("바나나가 사는곳 %s\n", m["바나나"])

	//////////////////////////////////////////////
	m := make(map[int]Product)

	m[10] = Product{"가습기", 500}
	m[50] = Product{"스타일러", 1500}
	m[161] = Product{"냉장고", 200}
	m[77] = Product{"세톽기", 1700}
	m[99] = Product{"건조기", 2800}

	for k, v := range m {
		fmt.Println(k, v)
	}

	m1 := make(map[int]int)

	m1[1] = 0
	m1[2] = 2
	m1[3] = 3
	delete(m1, 3)
	delete(m1, 4)
	fmt.Println(m1[3])
	fmt.Println(m1[2])

}
