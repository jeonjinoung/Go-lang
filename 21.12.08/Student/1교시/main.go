package main

import "fmt"

//인터페이스
type Stringer interface {
	String() string
}
//구조체
type Student struct{
	Age int
}
//스튜던트 타입의 String메서드
func(s *Student) String() string{
	return fmt.Sprintf("나이:%d",s.Age)
}

func PrintAge(stringer Stringer){
	s:=stringer.(*Student)//*Student type으로 type변환
	fmt.Printf("나이:%d\n", s.Age)
}

func main() {
	// 인터페이스 변수를 다른 구체화된 타입으로 변환할 수 있다.
	// 인터페이스를 본래 구체화된 타입으로 돌릴때
	//interface변수
	/*실행안함
	var a interface
	t:=a.(변경하려는 타입)
	*/

	/*실행함
	s := &Student{100}//*Student type변수 s
	PrintAge(s) //변수 S를 인터페이스 인수로 PrintAge()부름
	*/

	/*실행안함
	var stringer Stringer
	stringer.(*Student)//역참조 //인터페이스를 구현하고있지않기때문 스튜던트타입으로 변환이 불필요하다
	*/

}

/*
1. interface는 Method 집합
2. interface에서 정의한 Method 집합을 가진 모든 type은 interface로 쓸수 있다.
3. 모든 type이 빈 interface 변수값을 쓸 수 있다.
4. interface 변환을 사용하려면 interface 변수를 구체화된 type이나 다른 interface로 변경가능
 */