package main

import "fmt"

type Product struct{
	Name string
	value int
}
func main() {

	//key와 value로 어떤type 넣을 수 있다.
	m := make(map[int]Product)
	m[10] = Product{"가습기", 500}
	m[50] = Product{"스타일러", 1500}
	m[161] = Product{"냉장고", 200}
	m[77] = Product{"세탁기", 1700}
	m[99] = Product{"건조기", 2800}


	//range키워드를 사용하여 map을 돌리는데
	//배열이라 슬라이스 했을때는 
	
	//_ = amp v = value다 map에서는
	//for _, v := range v {
	for k, v := range m {
		fmt.Println(k, v)	
	}
	//입력순서 값과 상관없이 내용을 꼴아박아버리고 그데이터를 보관한다.
	//map은 속도가 굉장히 빠르다.너무빠르다

	m1:=make(map[int]int)
	m1[1] = 0
	m1[2] = 2
	m1[3] = 3

	delete(m1,3)
	delete(m1,4)
	fmt.Println(m1[3])
	fmt.Println(m1[4])
}
