package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b //함수 리터럴을 이용해서 더하기 함수 정의하고 리턴
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b //함수 리터럴을 이용해서 더하기 함수 정의하고 리턴
		}
	} else {
		return nil
	}
}

func main() {
	//함수리터럴 -> 이름이 없는 함수 -> 함수명을 쓰지않고 함수 type 변수값으로 대입되는 함수값
	//함수명이 없기 때문에 직접호출 x, 함수 타입변수로만 선언

	fn := getOperator("*")
	res := fn(3, 4)//function type 변수를 사용해서 function 호출
	fmt.Println(res)

	fn1 := getOperator("+")
	res1 := fn1(3, 4)//function type 변수를 사용해서 function 호출
	fmt.Println(res1)
	//////////////////////////////////////////////////////////////////////////////

	a := 0
	f:=func ()  {
		a += 10
	}
	a++
	f()//function type 변수
	fmt.Println(a)


}

/*
function type 변수
함수 리터럴
*/