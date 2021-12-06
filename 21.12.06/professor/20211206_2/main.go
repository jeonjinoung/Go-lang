package main

import "fmt"

func changeArr(arr [5]int) {
	arr[2] = 200
}
func changeSlice(slice []int) {
	slice[2] = 200
}
func main() {

	// type SliceHeader struct {
	// 	Data uintptr //실제 배열을 가리키는 포인터
	// 	Len  int     //요소 개수
	// 	Cap  int     //실제 배열의 길이
	// }

	//var slice = make([]int, 3, 5)

	// array := [5]int{1, 2, 3, 4, 5}
	// sli := []int{1, 2, 3, 4, 5}

	// changeArr(array)
	// changeSlice(sli)
	// fmt.Println(array)
	// fmt.Println(sli)

	// //slice2 := append(slice, 3, 5)

	// slice1 := make([]int, 3, 5) //len:3 cap:5 슬라이스 만듬
	// slice2 := append(slice1, 4, 5)
	// fmt.Println("슬라이스 1 : ", slice1, len(slice1), cap(slice1))
	// fmt.Println("슬라이스 2 : ", slice2, len(slice2), cap(slice2))
	/*
		만약 빈공간이 없으면 더 큰 배열을 만듬. 일반적으로 기존 배열의 2배
		->기존 배열의 요소를 새로운 배열에 복사->새로운 배열 맨뒤에 새 값을 추가
		cap : 새로운 배열의 길이 값
		len : 기존길이에 추가한 개수만큼 더한값
		포인터 : 새로운 배열을 가리키는 슬라이스 구조체를 return
	*/

	// slice3 := []int{1, 2, 3}       //len : 3, cap : 3
	// slice4 := append(slice3, 4, 5) //길이가 6개짜리 새로운 배열 만들고 슬라이스 배열의 모든값을 복사->맨뒤에 4,5 len : 5,cap : 6

	// fmt.Println("슬라이스 3 : ", slice3, len(slice3), cap(slice3))
	// fmt.Println("슬라이스 4 : ", slice4, len(slice4), cap(slice4))

	// slice3[1] = 200
	// fmt.Println("슬라이스 3 : ", slice3, len(slice3), cap(slice3))

	//슬라이싱  : 배열의 일부를 집어낸다
	//array[starIndex : endIndex]
	//주의할점  : endIndex-1 까지
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:2] //슬라이싱

	fmt.Println("arr : ", arr)
	fmt.Println("slice : ", slice, len(slice), cap(slice))

	arr[1] = 100 //arr의 두번째 값 변경

	fmt.Println("arr : ", arr)
	fmt.Println("slice : ", slice, len(slice), cap(slice))

	slice = append(slice, 200)

	fmt.Println("arr : ", arr)
	fmt.Println("slice : ", slice, len(slice), cap(slice))

	///////////////////////////////
	//슬라이스를 슬라이싱
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := slice1[1:2]	//[2]
	//처음부터
	//slice2 := slice1[0:3]	//[1,2,3]
	//slice2 := slice1[:3]	//[1,2,3]

	//끝까지
	//slice2 := slice1[2:len(slice1)]	//[3,4,5]
	//slice2 := slice1[2:]	//[3,4,5]
	//전체->배열 전체를 가리키는 슬라이스 만들때 사용
	//slice2 := slice1[:]	//[3,4,5]

	//cap크기 조절하기(인덱스 3개로)
	//slice[startIndex:endIndex:maxIndex]

	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := slice1[1:3:4] //[2,3] cap : 3(maxIndex-startIndex)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1)) //슬라이스1이랑 같은길이의 슬라이스 만듬

	//슬라이스 1의 모든 값을 복사
	for i, v := range slice1 {
		slice2[i] = v
	}
	slice1[1] = 200
	fmt.Println(slice1)
	fmt.Println(slice2)

	//copy 함수

	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := make([]int, 3, 10) //len:3 cap : 10
	slice5 := make([]int, 10)    //len:10 cap : 10
	//첫번째 인수 : 복사한 결과를 저장하는 슬라이스 변수
	//두번째 인수 : 복사대상이 되는 슬라이스 변수
	copy1 := copy(slice4, slice3)
	copy2 := copy(slice5, slice3)

	fmt.Println(copy1, slice4)
	fmt.Println(copy2, slice5)

}
