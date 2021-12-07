package main

import "fmt"

type account struct {
	balance    int
	firstName  string
	secondName string
}

//포인터 메서드
func (a1 *account) PointerMethod(amount int) {
	a1.balance -= amount
}

//값 타입 메서드
func (a2 account) VaulteMethod(amount int) {
	a2.balance -= amount
}

func (a3 account) ReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	//리시버 두가지 타입으로 정의할 수 있다.
	//1. value type 2.pointer type

	var A *account = &account{200, "홍길동", "이길동"}
	A.PointerMethod(10)//pointermethod 불러와서 
	fmt.Println(A.balance)//그 값을 출력하는것 190

	//VaulteMethod의 A와 PointerMethod A는 다르다
	//왜냐하면 valteMethod는 PointerMethod의 값을 가져오기때문에 메모리 주소가 다르다
	A.VaulteMethod(20)//VaulteMethod
	fmt.Println(A.balance)//190 

	var B account = A.ReturnValue(30)
	fmt.Println(B.balance)//160
	
	B.PointerMethod(10)
	fmt.Println(B.balance)//150

}
/*

pointer와 value의 차이
pointer의 가르키면 메모리의 주소값을 가르키게된다.
value타입을 오출을하면 resiver타입의 모든값이 복사가된다. -> 구조체의 모든값이 복사가된다는 내용이다.

리시버가 있으면 메서드고 없으면 일반 함수이다.
리시버는 메서드를 호출하는 주체다 / 메서드는 리시버를 통해서만 호출
메서드는 리시버에 속한 기능이다.
포인터 메서드는 인스턴스 중심. 값을 변경 할 수 있다.
값 타입은 값이 모두 복사. 인스턴스에 접근할 수 없어 복사하는 양에 따라서 성능 이슈가 발생할 수 있다.


*/