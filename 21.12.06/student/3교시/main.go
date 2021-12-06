package main

import "fmt"

func changeArr(arr [5]int) {
	arr[2] = 200
}

func changeSlice(slice []int) {
	slice[2] = 200
}

func main() {

	//make함수를 써서 함수를 만들때는 인수를 2개또는 3개를 넣는다.
	//var slice = make([]int, 3, 5)

	// array := [5]int{1, 2, 3, 4, 5}// 여기에서 지금 40byte 
	// sli := []int{1, 2, 3, 4, 5}
	
	// changeArr(array)
	// //40byte가 복사된다는것이고 배열의 모든값이 복사된다.{1, 2, 3, 4, 5}
	
	// changeSlice(sli)
	// fmt.Println(array)
	// fmt.Println(sli)

	// //slice2 := append(slice,4,5)

	// slice1 := make([]int, 3,5 )//len : 3 cap :5 슬라이스 만듬
	// slice2 := append(slice1, 4, 5,)
	// fmt.Println("슬라이스 1 : ", slice1, len(slice1),cap(slice1))
	// fmt.Println("슬라이스 2 : ", slice2, len(slice2),cap(slice2))
	/*만약 빈공간이 없으면 더 큰 배열을 만듬 . 일반적으로 기존 배열의 2배
	->기존 배열의 요소를 새로운 배열에 복사-> 새로운 배열 맨뒤에 새 값을 추가
	cap : 새로운 배열의 길이 값
	len : 기존 길이에 추가한 개수만큼 더한값
	pointer : 새로운 배열을 가르키는 슬라이스 구조체를 return
	*/

	slice3 := []int{1,2,3} //len : 3 cap : 3
	slice4 := append(slice3, 4, 5)
	//길이가 6개짜리의 새로운 배열을 만든다음에 슬라이스의 모든값을 복사를하고 맨뒤에 4,5를 추가해서 리턴해주는것

	fmt.Println("슬라이스 3", slice3, len(slice3), cap(slice3))
	fmt.Println("슬라이스 4", slice4, len(slice4), cap(slice4))
	
	slice3[1] = 200
	fmt.Println("슬라이스 3", slice3, len(slice3), cap(slice3))

	//슬라이싱 : 배열의 일부를 집어낸다.
	//array[starIndex : endIndex] 
	//주의할점 : endIndex-1 까지

	arr := [5]int{1,2,3,4,5}
	slice:= arr[1:2]//slicing 슬라이싱 

	fmt.Println("arr : ", arr)
	fmt.Println("slice : ", slice, len(slice), cap(slice))

	arr[1] = 100//arr의 두번째 값 변경

	fmt.Println("arr : ", arr)
	fmt.Println("slice : ", slice, len(slice), cap(slice))

	slice = append(slice,200)
	fmt.Println("arr : ", arr)
	fmt.Println("slice : ", slice, len(slice), cap(slice))

	///////////////////////////////////////////////////////
	//슬라이스를 슬라이싱
	// slice1 := []int{1,2,3,4,5}
	// slice2 := slice1[1:2]

	//cap크기 조절하기(인덱스 3개로)
	//slice[startIndex:endIndex:maxIndex]

	// slice1 := []int{1,2,3,4,5}
	// slcie2 := slice1[1:2:4] //[2,3] maxIndex에서 startIndex-1을해준다. cap : 3(maxIndex-startIndex)

	slice1 := []int{1,2,3,4,5}
	slice2 := make([]int,len(slice1))//slice1이랑 같은 len의 slice 만듬

	//슬라이스 1의 모든값을 복사
	for i, v := range slice1 {
		slice2[i] = v
	}
	slice1[1] = 200
	fmt.Println(slice1)
	fmt.Println(slice2)


	//copy함수
	slice5 := []int{1,2,3,4,5}
	slice6 := make([]int, 3, 10) //len : 3 cap : 10
	slice7 := make([]int, 10) //len : 10 cap : 10
	//첫번째 인수 : 복사한 결과를 저장하는 slice6 인수
	//두번째 인수 : 복사 대상이 되는 slice5 변수
	copy1 := copy(slice6,slice5)
	copy2 := copy(slice7,slice5)

	fmt.Println(copy1, slice6)
	fmt.Println(copy2, slice7)


}
// 모든변수의 이동은 복사를 베이스로 하게 된다. 포인터는 포인터의 값이 메모리주소가되고
// 구조체가 복사될때는 구조체의 안에잇는 모드녀석들이 복사
// 배열은 배열의 모든값이 복사가되는것
