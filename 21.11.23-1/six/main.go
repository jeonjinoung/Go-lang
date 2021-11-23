package main

import "fmt"

func main() {
	var num int = 5
	var pnum = &num
	
	fmt.Println("num : ", num)   //num 값
	// num :  5
	fmt.Println("pnum :", pnum)  //num의 메모리 주소
	// pnum : 0xc000014088
	fmt.Println("pnum :", *pnum) //num의 주소로 메모리에 할당돼있는 값 접근
	// pnum : 5

	*pnum++
	fmt.Println("num : ", num)
	// num :  6
	fmt.Println("pnum :", *pnum)
	// pnum : 6
	//포인터 연산자를 이용한 값 변경
}
