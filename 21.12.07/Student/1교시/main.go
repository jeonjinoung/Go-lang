package main

import "fmt"

//패키지 내에서 type 키워드로 선언된 = 로컬타입 => 모두 리시버로 사용한다.
type account struct{
	b int 
}

//일반함수 표현
func Function(a * account, amount int){
	a.b -= amount
	} 

	
//메서드로 표현하면?
	//리시버 = (a*account)
	//Method는 a*account안에 속한다.
func(a*account)Method(amount int ){
		a.b -= amount
	}
	
///////////////////////////////////////////////////////////////////////
	
//사용자 정의 타입 -> 리시버 타입이 될수 있기 때문에 -> 기본 내장 type도 별칭 type이 될수 있기때문에 매서드가 될 수 있다.
type myNum int

//이 별칭(myNum) type을 리시버로 받는다.
func(a myNum)add(b int)int {
	return int(a) + b
}

func main() {

	//메서드 선언하기
	//리시버 부분 찾기 -> (r Rect) 이부분이 리시버
	//메서드 이름 찾기 ->  cal() 메서드의 이름
	// r = r.width r.height 똑같은 메개변수로 사용된다.

	// func (r Rect) cal()int{
	// 	return r.width* r.height
	// }
	
	test := &account{200} // b가 200
	Function(test,30)

	//.을 찍어서 내부 메서드에 접근한다.
	//test.b
	//test.Method()
	test.Method(20)
	fmt.Printf("%d\n", test.b)

///////////////////////////////////////////////////
//타입 형변환을 지원하기 때문에 형변환을 때린것이다.

	//
	var a myNum = 10
	fmt.Println(a.add(20))

	// int type 
	var b int = 20
	fmt.Println(myNum(b).add(100))
}
/*
method : 함수의 일종임. go에서는 class가 없다 -> 구조체 밖에 매서드를 지정해줘야한다.
resever(리시버) : 메서드가 속하는 타입을 알려주는 녀석-> 어느 구조체에 속하냐? 표시를 해주는것 -> 리시버사용
매서드를 좀 볼게요 함수의 일종입니다. 문법이 불편하다 생각한다.
이유는 : C++에서는 클래스가 있는데 맴버함수를 만들수가 있다 소속으로

모든 로컬타입이 리시버 // 별칭타입 리시버 // 인티져 같은 내장타입들도 리시버를 통해서 메서드를 가질수 있다.

///////////////////////////////////////////////////////////////////////////////////////////////////
형변환
//매개변수하나를 함수로 바꾸었을뿐?? 중요한차이가 있다.?
//어디소속인것이냐 그것이 중요한것이다.

Go는 객체지향프로그램이 아니다 해석의 여지에따라 틀리다.
접근을 어떻게하느냐에 따라 틀리다.
매서드 인터페이스를 지원한다.->지향을 하는거지
*/