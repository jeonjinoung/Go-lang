package main

import "fmt"

//Go lang에서는 권장사항으로 er을 붙여줘야한다.
type Tester interface{
	String() string
}

type Student struct{
	Name string
	Age int
}
//덕 타이핑 지원하지 않는 언어들은 Test 명시해줘야한다.
// type Student struct implements Test
//Stuent의 구조체의 String이라는 메서드 

/*
String 메서드를 포함하고 있다. 그렇기 때문에 s Student의 타입은 Tester interface로 사용될 수 있다.

** Sprinf 서식을 만들어서 문자열로 만들어주는 녀섴
*/
func(s Student) String() string{	
	return fmt.Sprintf("하이 %d, %s", s.Age, s.Name)
}

//빈 인터페이스
//빈 인터페이스를 인수로 받는 함수이다.
func PrintValue(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("int %d\n", int(t))
	case float64:
		fmt.Printf("int %f\n", float64(t))
	case string:
		fmt.Printf("int %s\n", string(t))
	default:
		fmt.Printf("위에 없는 다른 타입 %T:%v\n", t,t)
	}
}

type People struct{
	Age int
}

type Mover interface {
	Moving()
}
func main() {
	// 인터페이스도 타입중 하나임
	// 인터페이스 변수 선언 가능. 변수의 값으로도 쓸수 있다.
	// 인터페이스(interface) : 구현을 포함하지 않는 메서드 집합
	// type DuckInterface interface{
		//메서드의 집합이 들어가겠네요
		//주의할점 : 메서드는 반드시 메서드명이 있어야한다.
		//매개변수와 return이 다르더라도 이름이 같은 메서드는 있을수 없다
		//interface에서는 메서드 구현을 포함하지 않는다.
	// 	walk()
	// 	Move{distans int } int 
	// }

	// type Tester interface{
	// 	String() string
		//String(int) string //함수 이름이 같아서
		//_(a int) // 메서드 이름이 없어서
	// }

	st :=Student{"이순신", 10} //Student  type
	
	var t Tester // Tester type

	t = st // t 값으로 st를 대입을 한것이다.

	fmt.Printf("%s\n", t.String())
	
	//덕 타이핑
	//Type선언 ->interface 구현 여부를 명시적으로 나타낼 필요없이
	//interface에 정의한 Method 포함 여부만으로 결정하는 방식

	//인터페이스 또 다른 기능
	//1.인터페이스 기본값
	//2.빈 인터페이스
	//3.포함된 인터페이스

	////////////////////////////////////////////////////////////////////////////////

	//구현되지않은 포함한 인터페이스
	// type Reader interface{
	// 	Read()(n int, err error)
	// 	Close()error
	// }
	
	// type Writer interface{
	// 	Write()(n int, err error)
	// 	Close()error
	// }

		///////////////////

	// //인터페이스는 또다른 인터페이스를 포함할 수 있다.
	// type ReadWriter interface{
	// 	Reader
	// 	Writer
	// }

	////////////////////////////////////////////////////////////

	//빈 인터페이스
	//Method를 들고있지 않은 빈 녀석

	PrintValue(10)
	PrintValue(3.141592)
	PrintValue("빈 인터페이스")
	PrintValue(People{100})

	//유효하지않는 메모리 주소값 실행중에 발생한에러 런타임애러
	//인터페이스 값이 항상 닐이 아닌지 확인을해주어야한다.
	var m Mover
	m.Moving()
}
/*

인터페이스(interface) : 구현을 포함하지 않는 메서드 집합
구체화 된 타입이 아닌 인터페이스만 가지고 메서드를 호출 -> 프로그램 요구사항 변경시 유연하게 대처 가능
Go에서는 인터페이스 구현 여부를 그타입이 인터페이스에 해당하는 메서드를 가지고 있는지로
판단하는 덕 타이핑을 지원한다. 
type => interface name 써주고 -> interface keword


*/