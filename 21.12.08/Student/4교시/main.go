package main

import "fmt"

func cap1() {
	f := make([]func(), 3) //function 리터럴 3개를 가진 slice
	fmt.Println("캡1")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)//capture가 된 아이를 출력
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func cap2() {
	f := make([]func(), 3) //function 리터럴 3개를 가진 slice 서로다른 부위변수를 캡처를 뜨는것이다.
	fmt.Println("캡2")
	for i := 0; i < 3; i++ {
		v := i //v변수에 i value 값을 복사 v는 local 변수 지역변수 새로운 변수
		f[i] = func(){
			fmt.Println(v)
		} 
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}


func main() {
	//함수 리터럴 외부 변수를 내부로 가져오는 거를 캡처(capture)라고한다.
	//capture는 value copy가 아니라 레퍼런스(참조) 형태로 가져옴
	//변수의 주소를 포인터 값으로 복사하는것이라고 보면 된다.

	cap1()
	cap2()

}

/*

 */