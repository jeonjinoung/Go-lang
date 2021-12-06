package main

import "fmt"

func main() {
	//슬라이스 : 동적배열->자동으로 배열 크기를 증가시켜준다.
	//길이가 요소개수에 따라 자동으로 증가해 관리가 편함.
	//슬라이싱 기능을 사용해서 배열의 일부를 나타내는 슬라이스를 만들수 있다.

	//var arr[10]int->10개까지...

	//슬라이스를 초기화 하지 않으면 길이가 0인 슬라이스가 만들어진다.
	var slice []int

	if len(slice) == 0 {
		fmt.Println("슬라이스가 비어 있다", slice)
	}
	//slice[1] = 10 //할당되지 않은 메모리공간에 접근->에러발생
	//fmt.Println(slice)

	//{}를 사용한 초기화
	//var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}
	fmt.Println(slice2)

	//var arr = [...]int{1,2,3}//배열
	//var slice = []int{1,2,3}//슬라이스

	//make()를 이용한 초기화
	//var slice3 = make([]int, 3)//첫번째 인수로 타입,두번째 인수는 길이

	//슬라이스 접근 : 배열과 같음
	// var slice4 = make([]int, 3)
	// slice4[1] = 5 //[0,5,0]

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

	//요소추가
	var slice6 = []int{1, 2, 3}

	//첫번째 인수 :  추가하려는 슬라이스, 두번째 인수에는 추가하려는 요소
	slice7 := append(slice6, 4, 5, 6, 7, 8) //추가

	fmt.Println(slice6)
	fmt.Println(slice7)
	//append : 첫번째 인수로 들어온 슬라이스의 값을 변경하는게 아니라 요소가 추가된 새로운 슬라이스를 반환
	//기존 슬라이스에 요소를 추가하고 싶다? append()결과를 기존 슬라이스에 대입해서 변경해야 된다.

	var slice8 []int
	//포문으로 하나씩 추가
	for i := 1; i <= 10; i++ {
		slice8 = append(slice8, i)
	}
	slice8 = append(slice8, 11, 3, 315, 123)
	fmt.Println(slice8)

}
