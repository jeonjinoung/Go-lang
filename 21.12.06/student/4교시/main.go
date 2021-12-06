package main

import (
	"fmt"
	"sort"
)

// type Student struct{
// 	Name string
// 	Age int
// }

// type Students []Student

func main() {

	// slice := []int{1, 2, 3, 4, 5, 6}
	// index := 2
	// //네번째값부터 하나씩 땡겨준다고 생각하면됨
	// for i := index + 1; i < len(slice); i++ {
	// 	slice[i-1] = slice[i]
	// }
	// //마지막값을 짤라준거라 보면됨
	// slice = slice[:len(slice)-1]
	// //그러면 값이 = 1, 2, 4, 5, 6 프린트됨
	// fmt.Println(slice)

	//삽입하려면??
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 0)
	index := 2
	//맨뒤부터 추가하려는 위치값을 하나씩 뒤로 옮긴다.
	for i := len(slice) - 2; i >= index; i-- {
		slice[i+1] = slice[i]
	}
	slice[index] = 299
	fmt.Println(slice)

	//작은수부터 앞으로 나열하는 함수 sort
	s := []int{5,100,3,1,7}
	sort.Ints(s)
	fmt.Println(s)
 
}