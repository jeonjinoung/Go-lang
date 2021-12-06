package main

import "fmt"

func main() {
	//슬라이스 : 동적배열이다. -> 자동으로 배열 크기를 증가시켜준다.
	//장점 : 길이가 요소 개수에따라 자동으로 증가해 관리가 편함
	//슬라이싱 기능을 사용해서 배열의 일부를 나타내는 슬라이스를 만들수 있다.

	//var arr[10]int -> 10개까지...

	//슬라이스를 초기화 하지 않으면 길이가 0인 슬라이스가 만들어진다.
	var slice []int

	if len(slice) == 0 {
		fmt.Println("슬라이스가 비어있다.", slice)
	}

	// slice[1] = 10 //할당되지 않은 메모리 공간에 접근 -> 애러 발생 panic
	// fmt.Println(slice)
	
	//초기화 방법
	//var slice1 = []int{1,2,3} -> {}를 이용한 초기화 방법

	var slice2 = []int {1, 5:2, 10:3}
	fmt.Println(slice2)

	//var arr = [...]int{1,2,3}//배열 
	//var slice = []int{1,2,3}//슬라이스 

	//make()를 이용한 초기화 
	//길이가 3인 int슬라이스 값을 같는다.
	//첫번째 인수로 타입을 넣어주고 두번째 인수로는 길이(len)넣어준다.
	//var slice3 = make([]int, 3)

	//슬라이스 접근 : 배열과 같음
	// var slice4 = make([]int, 3)
	// slice4[1] = 5//[0,5,0]

	var slice5 = make([]int, 5)
	
	//슬라이스 순회->동적으로 길이가 늘어나는거 말고 배열과 동일함.
	//초기화
	for i := 0; i < len(slice5); i++ {
		slice5[i] = i + 1
	}
	//출력
	for _, v := range slice5 {
		fmt.Println(v)
	}


	//요소 추가

	var slice6 = []int{1,2,3}
	//append라는 함수를 통해 추가해줄수 있다.
	//첫번째 인수 : 추가하려는 슬라이스, 두번째는 추가하려는 요소
	//slice7 := append(slice6,4) slice6는 추가하려는 슬라이스
	//slice7 := append(slice6,4) 4는 추가하려는 요소

	slice7 := append(slice6,4,4,5,6,5,3,5)
	fmt.Println(slice6)
	fmt.Println(slice7)

	var slice8 []int
	for i := 1; i <= 10; i++ {
		slice8 = append(slice8, i)
	}
	slice8 = append(slice8, 11, 3, 315, 123)
	fmt.Println(slice8)
}


/*
Slice에대해 알아볼게 자료구조야
할당되지않은 메모리공간에 값을 넣을때 panic현상 발생
배열이랑 사용법이 똑같다. 
배열 요소를 추가해서 길이를 늘릴수가 있다. 내장함수 append

vactor push.back 때리면 뒤에서부터 데이터가 들어간다.
여러개의 여러가지의 데이터를 넣을수 있다는것이다
*/