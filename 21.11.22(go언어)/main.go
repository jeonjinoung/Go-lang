package main

//페키지를 여러개 쓸때는 소괄호로 열어주면된다.
import (
	"fmt"
	"unsafe"
)

/*

페키지선언 : 이코드가 어떤 페키지에 속하냐?
go언어의 모든 코드는 반드시 패키지 선언으로 시작
packge main : main 패키지에 속한 코드다!! 컴파일러한테 알려줌
go언어는 패키지로 시작하고 package main은 프로그램 시작점이 있는 페키지다.!!!


*/


func main()  {

	//printin 엔터치는거와같다.
	
	fmt.Print("나는 프린트함수")
	fmt.Println("Hello Wolrd")
	fmt.Printf("나는 프린트함수\n")


	
	//낭비다 5만 표현하는것인데
	var num = 5
	fmt.Println("변수 num의 사이즈 ", unsafe.Sizeof(num))

	//선언대입문 : :=를 써서 var키워드, type을 생략
	num2 :=10
	fmt.Println("키워드와 타입생략 방법", num2)


	var a int = 10;
	var b int = 20;
	var msg string = "hellow"
	
	//변수는 항상 동일하게 안그러면 오류가 난다.
	//var b float32 = 20.00;
	//var plus int = a + b


	var plus int = a + b
	fmt.Println(a)
	fmt.Println(msg)
	fmt.Println(plus)

	//첫번째
	//타입의 필요성 -> 공간의 크기 때문에
	//변수는 변수라는녀섴은 메모리주소를 가르킨다. 그런데 메모리주소는 값이 있는 10이있는 값이라는 시작주소만을 알려준다.


	//두번째
	//타입을 알아야지만 컴퓨터가 데이터를 해석할 수 있다. -> 제공하는 숫자타입이 여러가지있는데 

	//변수속성은 내가지를 씀 -> 이름, 값 , 타입 , 주소
	//타입은 많은데 ->
	/*
	
	

	*/
	//uint8 부호가 없는 정수
	//uint16 0~65535까지 표현가능
	//2바이트인데 부호없는 정수
	//	unit32
	//	unit64

	//int -> 32bite 운영체졔면 int32
	//int -> 64bite 운영체졔면 int64



	//비트 0000 0000 uint8 u 언사인드 -> 크기는 8비트니까 1바이트가되겠네용 -> 0에서 255까지 표현이 가능하다.
 		//var abc uint8 = 255		 
		//var abc uint8 = 256 0~255까지이기때문이다.

	//Sizeof : 크기를 측정해주는 녀석  ->

	var int_size int
	var int_size8 int8
	var int_size16 int16
	var int_size32 int32
	var int_size64 int64
	var uint_size64 int64
	//"int형 사이즈 : " 8바이트 나온다.

	/*
	//실수형
	float32 
	float64

	//복소수형
	complex64
	complex128
	
	byte(별칭) = uint8

	rune : 문자하나를 나타낼때 사용함 -> UTF-8문자 하나를 나타낼때 사용
	rune(별칭) = int32 
	bool = 참과 거짓
	string = 문자열

	*/
	

	fmt.Println("int형 사이즈 : ", unsafe.Sizeof(int_size), "바이트")
	fmt.Println("int8형 사이즈 : ", unsafe.Sizeof(int_size8), "바이트")
	fmt.Println("int16형 사이즈 : ", unsafe.Sizeof(int_size16), "바이트")
	fmt.Println("int32형 사이즈 : ", unsafe.Sizeof(int_size32), "바이트")
	fmt.Println("int64형 사이즈 : ", unsafe.Sizeof(int_size64), "바이트")
	
	//내가 표현하고 싶은 범위는 
	fmt.Println("uint64형 사이즈 : ", unsafe.Sizeof(uint_size64), "바이트")
	
	
	/*
	if true{

	}

	for i := 0; i < count; i++ {
		
	}
	*/

}

//기계어로 변환해놨다가 사용하는언어를 정적컴파일러 언어?

//exe파일 -> 정적컴파일러 -> 실행파일 -> 기계어코드

//실행할때 변환과정이 필요가 없다. 

//type애러를 컴파일러 시점에서 발견한다는것은 안전성이 뛰어나다.

//동적 컴파일러언어 runtime의 기계어로 변환되는 언어?

//Go 언어는 정적 컴파일러언어!!

//어떤 환경 플렛폼에 맞는 exe파일을 따로 만들어주어야한다.

//내부환경 변수만 가지고와가지고 여러가지 플렛폼을 맞도록 exe 만들수가있다.

//프로그램언어를 나눌때 type검사를 강하게하는 녀섴 유도리있게하는녀석



//type검사를 강하게 하지않는다. 약하게한다 -> 자바스크립

//type검사를 강하게 한다 -> Go 언어


//강type언어냐 약type언어냐 


/*
Go 언어는 가비쥐 컬렉터 -> 쓰레기수집가 -> 메모리 회수를 알아서 한다. -> 2009년에 발달함 -> 프로그램 언어순위에 Top 3안에 든다.
가비지 컬렉터를 사용하는 언어 -> 메모리 운영면에서 좋다. -> CPU를 사용한다.(단점) -> Java, Go언어
Go 자바랑 파이썬이랑 장점을 끌어모아두었다. 
상속 지원이 안됨
메서드 있음
인터페이스라는 녀섴이 있음
익명함수를 지원합니다.
포인터 있음
제네릭 프로그래밍 x
네임스페이스 x 페키지 단위로 구분
*/


/*
C언어

Class라는 분법을 지원함


/*
class main
{
private:

public:
main()
};
*/


/*
int main(void)  {
	//가비지 컬렉터를 가지고 있지 않는 언어는 할당 후 제거를 해주어야 한다. -> CPU를 사용안한다. -> C언어...

	new <--라는 녀석으로 동적 할당을 한다. <- 메모리공간 이 얼마정도 필요하다 빌려줘라 <- 운영체제에서 빌려준다. <-약간 느리다.

	//else if  빌렸는데 못값는다면 메모리누수 -> 메모리가 줄줄센다.? -> 메모리 점유율이 점점 높아진다. -> 이것을 보안하기위해 나온녀섴이 가비쥐 컬렉터
	스텍영역 -> 매개변수 -> 
	힘영역 -> 동적할당은 힘영역  -> 
}
*/
