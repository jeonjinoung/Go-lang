package main

import (
	"fmt"
	"unsafe"
)

//메인함수 밖에 있는것은 전역변수이다.

func main(){
	fmt.Println("Hellow World")
	
	var int_size int

	fmt.Println("int형 사이즈 :", unsafe.Sizeof(int_size),"바이트")


	// Javascritp에서는 되던게 안됨
	// var num = 1
	// num = 'string'


	// 단일문자열은 싱글 ''
	// ""

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	
	//타입변환(형변환)

	// 타입이 안맞으면 애초부터 오류를 방지해줌 -> 명확하다.
	// num3:= 3
	// var fnum float64= 3.5
	
	// var c int = int(fnum)//타입이 틀리기 때문에 대입도 되지 않는다.
	
	// fmt.Println(c)// -> 3이라는 값이 나온다. -> 3.5에서 -> int 정수로 표현하여서 소수점 자리가 날라간다.

	// res:= num3* fnum//타입이 다르기 때문에 연산이 되지 않는다.

	// var e int64 = 7

	// res1: num3 * e//같은 정수지만 타입이 틀리기때문에 계산이 안된다. num3와 num3가 같은 정수(int)지만 type이 다르다.

	//그렇다면 이것을 사용하려면 형변환 -> 타입변환을 해야한다.

	//큰 범위 에서 작은 범위로 바꾸면 값이 바뀌어진다.

	


	/*
	
	00000000 00000000 00000001 01101000 32bit 4byte => 360
	01101000
	64,32,16,8,4,2,1 -> 8+32+64 ->104
	00000001
	128
	*/

	var BigType int32 = 360
	var smalltype int8 = int8(BigType)

	fmt.Println(smalltype)

	

///////////////////////////////////////////////////////////////////////////////////////////////////



	//지역변수 와 전역변수에 대한 개념
	{
		//지역변수
		var s int = 50
		fmt.Println(s)
	}
	// fmt.Println(s)

	var defaultValueInteager int8
	var defaultValuefloat int32
	var isChack bool
	var str string

	//초기값을 넣지 않으면 아래와 같이 디폴트값이 들어감
	//그외...nil : 정의되지않은 메모리주소를 나타내지않을때
	fmt.Println("초기값을 넣지않으면 ? 디폴드값이 슈슛~ : ", defaultValueInteager)
	// 초기값을 넣지않으면 ? 디폴드값이 슈슛~ :  0
	fmt.Println("초기값을 넣지않으면 ? 디폴드값이 슈슛~ : ", defaultValuefloat)
	// 초기값을 넣지않으면 ? 디폴드값이 슈슛~ :  0	
	fmt.Println("초기값을 넣지않으면 ? 디폴드값이 슈슛~ : ", isChack)
	// 초기값을 넣지않으면 ? 디폴드값이 슈슛~ :  false
	fmt.Println("초기값을 넣지않으면 ? 디폴드값이 슈슛~ : ", str)
	// 초기값을 넣지않으면 ? 디폴드값이 슈슛~ :
	


}