package main

import "fmt"

func main() {

	// [sting] = key / sting = value
	//m:make(map[string(key)]string(value))
	//map은 key와 value 쌍으로 이루어져있다.
	m := make(map[string]string) //맵 생성
	m["이창희"] = "서울시 강남구"         //키와 값을 추가 
	m["바나나"] = "청송 교도소"          //키와 값을 추가
	m["홍길동"] = "한라산"             //키와 값을 추가
	m["김정은"] = "평양"              //키와 값을 추가

	m["바나나"] = "유치장으로 이동"              //value값을 key값을 재선언하여 변경
	//반환하는 key값으로 value 값을 찾아주는것
	fmt.Printf("바나나의 사는곳 %s\n", m["바나나"])
}

/*
map이란 자주사용하기에 go lang에 내장되어있다.

*/