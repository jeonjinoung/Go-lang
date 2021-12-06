package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func stringToUpper(str string)string{
	var a string
	for _,  v := range str {
		if v >= 'a' && v <= 'z' {
			a += string('A' + (v - 'a'))//
		}else{
			a += string(v)
		}
	}
	return a
}

func stringToUpper1(str string)string{
	var b strings.Builder
	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			b.WriteRune('A' + (v - 'a'))
		}else{
			b.WriteRune(v)
		}
	}
	return b.String()
}

func main() {
	//문자열 합산
	var str string = "Hello"
	//변수 스트링으로 선언 그값으로 "Hello"
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))//내부 구조체로 변경
	//reflect에대해 간단히 이야기하자면
	

	address := stringHeader.Data//DAta필드값을 변수로 저장
	str+="World"

	address1:=stringHeader.Data//데이터필드값을 변수로 저장
	str+="Hi"
	address2:=stringHeader.Data//데이터필드값을 변수로 저장

	fmt.Println(str)
	//HelloWorldHi
	fmt.Printf("주소 : \t%x\n", address)
	// 주소 : 	32e602
	fmt.Printf("주소 : \t%x\n", address1)
	// 주소 : 	c000014090
	fmt.Printf("주소 : \t%x\n", address2)
	// 주소 : 	c0000140a0

	//위에서 알 수 있듯이 기존 문자열 메모리공간을 건드리지않고 새로운 메모리 공간을 잡는다.
	//이말은 문자열 불변의 원칙을 지키게 된다.
	//문자열 합 연산이 빈번하게 발생되면 메모리가 낭비된다. 만약 빈번하게 사용된다고 하면 
	//string패키지의 Builder를 이용해 메모리 낭비를 줄이도록 한다.

	/////////////////////////////////////////////////////////////////////////////////////////////

	var upperString string = "Hello world"
	fmt.Println(stringToUpper(upperString))
	// HELLO WORLD
	fmt.Println(stringToUpper1(upperString))
	// HELLO WORLD
	/////////////////////////////////////////////////////////////////////////////////////////////

	fmt.Printf("주소 : \t%x\n", upperString)
	// 주소 : 	48656c6c6f20776f726c64
	fmt.Printf("주소 : \t%x\n", upperString)
	// 주소 : 	48656c6c6f20776f726c64

}

//주소가 각각 다르다 str 끼리 같은 str이지만 주소는 다르다
//이말은 뭐냐? 기존에 있는 문자열 메모리공간을 안건드린다.
//맨처음에 만든 메모리공간을 건드리지않고 새로운 공간을 만들어준다.
//저저번주에 했었떤 문자열 불변에 법칙에 준수한다.