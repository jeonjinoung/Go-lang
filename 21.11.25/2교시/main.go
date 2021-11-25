package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Age   int32
	Score float64
}

//bite의 크기는 더한값으로 = 9+ 9+ 1 = 19
//padding의 크기는 차이의 더한값으로 은 = 7+7+7 = 21 7씩 일어난다. 21+19 = 40
type Test struct{
	A int8 // 1
	B int  // 8
	C int8 // 1
	D int  // 8
	E int8 // 1
}
//bite의 크기는 더한값으로 = 3+16
//padding 크기는 차이의 더한값  8-1-1-1 = 5 ->  5(padding) + 3 + 16 (bite) => 24
type Test1 struct{
	//패딩이안일어난다.
	A int8 // 1
	C int8 // 1
	E int8 // 1
	//패딩이 안일어나고
	D int  // 8
	B int  // 8
}





//func PrintStudent(구조체)  {}
//func PrintStudent(st Student)  {}
func PrintStudent(st Student) {
	fmt.Printf("나이:%d, 점수: %.2f", st.Age, st.Score)
}

func main() {

	var s = Student{15, 20.5}
	// PrintStudent(s)
	// s1 := s // s 구조체 모든 멤버(필드)들이 s1복사된다.
	// PrintStudent(s1)//함수 호출할때도 구조체가 복사된다.

	//크기확인
	//메모리 정렬이라는 이유때문에 16이라는 값이나온다.
	//int = 8bite int 32 = 4bite int64 = 8bite

	
	//레지스터 크기가 4bite => 32bite  레지스터가 8bite => 64bite컴퓨터
	//컴퓨터가 레지스터와 크기가 똑같다면 효율적으로 데이터를 읽어들일수 있따. => 컴퓨터가 자체적으로 판단
	//64bite 컴퓨터에서 int 64 -> 8bite 면 시작주소가 100번지라면 -> 8의배수가 아니다 -> 레지스터 8에 맞춰서 정렬되있지않다
	//데이터 메모리를 읽어올때 성능상 손해를 입게된다. 처음부터 프로그래밍언어에서 데이터를 만들때 8의 주소인 메모리주소를 정렬한다.


	//결국엔 컴퓨터가 64bite 니까 성능저하를 막기위해 원래 4bite인 int32를 8bite로 여분있게 만들고 그리고 메모리주소를 8의 배수로 만들어 놓은거임
	fmt.Println(unsafe.Sizeof(s))
	//16bite
	var t =Test{1,2,3,4,5}
	fmt.Println(unsafe.Sizeof(t))
	//40bite
	var t1 =Test1{1,2,3,4,5}
	fmt.Println(unsafe.Sizeof(t1))
	//24bite
	
	


}
/*

구조체선언하면 구조체공간을 할당을 해줌
지금 이상황에서 보면 선언이되는순간 메모리공간을 할당을해주는데
int 8bite float64 8bite => 16bite Student 할당
구조체의 값을 복사할 수 있는데 

프로그램이 변화하면서 흐름대로 가다가 객체간 결합도 트렌드가 자꾸 바뀐다 어떻게하면 최대한 효율적으로 코드를 짜느냐
방향성이 객체간 결합도를 맞추고 연관있는데이터만 인직도???올리는방향??

함수라던가 배열이라던가 구조체 연직도??인직도?? 올려주는 역할을 한다.
함수는 당연히 관련코드를 묶어가지고 연직도를 높여가지고

배열은 같은타입의 응집도를 높이고 ....자시고 ...뭐시고 

*/