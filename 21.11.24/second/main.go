package main

import "fmt"

func SwitchGetAge() int  {
	return 30
}


func main() {

///////////////////////////////////////////////////////////////////////////////
//Go언어는 브레이크가 없어도 된다.
num:=3
	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	//3이 실횅된다.
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("에러")
	}
///////////////////////////////////////////////////////////////////////////////
//Go언어는 브레이크가 없어도 된다.
//하나이상의 값도 비교할 수 있다. 구분은 쉼표료 한다.

	day := "sunday"
	switch day {
	case "monday", "tuseday":
		fmt.Println("언제금요일이오냐?")
	case "wednsday","friday","sunday":
		fmt.Println("퇴근 시켜줘")
	//퇴근시켜줘 실행된다.
	}
///////////////////////////////////////////////////////////////////////////////
//case가 트루인 녀석을 찾는것이다. 비교할 값이 true

	num1 := 18
	//if문처럼 true가 되는 조건을 검사 할 수 있다.
	switch true {//비교할 값이 true -> case가 true인 녀석을 찾는다.
	case num1 < 10:
		fmt.Println("아니 크잖아??")
	case num1 > 30:
		fmt.Println("123123")
	//num1의 값이 10보다 크고 20보다 작다
	case num1 >= 10 && num1 < 20:
		fmt.Println("하하하하하하하")
	}
/////////////////////////////////////////////////////////////////////////////////////
//if문처럼 초기문을 넣을수 있다. 어디에 ??스위치에

	//SwitchGetAge의 리턴값이 30을 age와 비교한다는것이다.
	//switch age:=SwitchGetAge();age {
	//초기값 age:=SwitchGetAge
	//비교값(조건문) (리턴의 비교값과 비교);age

	switch age:=SwitchGetAge();age {
	case 10:
		fmt.Println("해당안됨")
	case 30:
		fmt.Println("여기가 맞아",age)
	}
/////////////////////////////////////////////////////////////////////////////////////
//진실판명조건문
//케이스가 true인것을 찾는데 return 값이 트루인녀석을 찾는것
	switch age:=SwitchGetAge(); true{
	case age<10:
		fmt.Println("가위")
	case age>30:
		fmt.Println("바위")
	case age>20 && age<=30:
		fmt.Println("보")
	}
//////////////////////////////////////////////////////////////////////////////////////
//fallthrough 다음케이스 까지 꺼내준다.!!

	abc := 3

	switch abc {
	case 1:
		fmt.Println("안녕")
	case 2:
		fmt.Println("안녕")
	case 3:
		fmt.Println("나는 3이다.")
		fallthrough
	case 4:
		fmt.Println("안녕12312312321")
	case 5:
		fmt.Println("안녕")
	}








}




/*
switch abc {
case condition:
}
*/