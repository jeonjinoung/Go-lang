package main

import (
	"fmt"
)

// type StringHeader struct{
// 	Data uintptr //문자열의 데이터가 있는 메모리주소를 나타내는 일종의 포인터
// 	Len int	//int타입의 길이
// }

func main() {
	str := "안녕 난 홍길동이야"
	str1 := str
	fmt.Println(str)
	fmt.Println()
	fmt.Println(str1)
	
	// StringHeader1 :=(*reflect.StringHeader)(unsafe.Pointer(&str))
	// StringHeader2 :=(*reflect.StringHeader)(unsafe.Pointer(&str1))
	// fmt.Println(StringHeader1)
	// fmt.Println(StringHeader2)

	var str2 string = "Hello World"
	//str2 = "경일아카데미"
	//fmt.Println(str2)
	// str2[3] = 'v'
	// []byte

	//->부호가 없는 정수타입의 가변길이 배열(1byte)이다.

	//type변환 
	var slice []byte = []byte(str2)
	slice[2] = 't'
	fmt.Println(str2)
	fmt.Printf("%s\n", slice)

}
//모든문자열은 1byte단위로 변환이 가능하다.


/*

Pointer또한 형변환이다. reflect.StringHeader를 변화하기위해 -> 강제로 변환시킨거임

//////////////////////////////////////////////////////////////////////

str
str1 둘다 구조체이다.
-> 구조체를 복사해올때 메모리값을 가져오는것인데
-> 구조체 크기만큼 복사가되는것이니까 2번이고 그값으로 Data uniptr, Len int값만 복사한다.
//str Data와 Len 값만 str1에 복사한다.




*/
