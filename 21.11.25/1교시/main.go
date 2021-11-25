package main

//구조체 : 여러 필드를 묶어서 사용한다.
//연관된 여러 데이터를 하나의 이름으로 묶어서 사용한다.

//1. type : 이 키워드를 통해 새로운 사용자 정의 타입을 정의한다.
//2. name : type명을 적는다. 타입명
//2.1 Name : 대문자면 퍼블릭과 프라이빗의 차이 // 프라이빗은 나만아느녀석 // 퍼블릭은 만인의 연인
//3. struct : 타입의 종류
//type name struct{

import "fmt"
	


type Student struct{
		Name string
		Age int
		Score float32
		Number int
	}

//타입안에 또다른 값을 넣을수 있다.?
type User struct{
	Name 	string
	ID 		string
	Password int
}
// 구조체 안에 구조체를 넣을 수 있다.
type HardUser struct{

	UserInfo User//위에 있는 User구조체
	Level int
	Age int
}


type HardUser1 struct{
//필드명을 생략하고하려면 필드명을 지워주면된다.
	User//위에 있는 User구조체
	Level int
	Age int
}
func main() {

	// var Student_name string
	// var Student_age int
	// var Student_phonNumber int


	var st Student
	st.Age = 15
	st.Number = 10
	st.Name = "캡틴 아메리카도"
	st.Score = 3.14

	fmt.Println(st)
	fmt.Println("학생이름: ", st.Name)
	fmt.Println("학생이름: ", st.Age)
	fmt.Println("학생이름: ", st.Number)
	fmt.Println("학생이름: ", st.Score)


	
	// 초기화 방법
	// 초기화 : 말그대로 양식은 그대로 쓰고 값만 변경하는 방법

	//한줄에 전부 쓸수 있다.
	var st1 Student = Student{"달마대사", 1000, 30.5, 20, }
	
	//여러줄에 쓴는방법

	var st2 Student = Student{
		"원효대사",
		2000,
		10.1,
		50}

	fmt.Println(st1)
	fmt.Println(st2)

	//일부만 초기화 하는법
	var st3 Student = Student{Name: "신사임당", Age: 1234, }
	fmt.Println(st3)
	
	
	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//구조체를 불러오는 방법
	var normalUser = User{"홍길동", "오늘만산다", 100}
	fmt.Println(normalUser)
	
	//구조체안에 구조체를 불러오는방법
	var HardUser = HardUser{User{"놀부", "흥부가 기가막혀", 200},10 ,20}
	fmt.Println(HardUser)
	


	////////////////////////////////////////////
	//구조체 필드명 생략하지 않은방법

	// %s String  / %d int
	fmt.Printf("과금 유저 이름 : %s 아이디 : %s 비밀번호 : %d 레벨 : %d 나이 : %d", 
	//구조체 안에 구조체를 접근하려면 점을 두번찍어야하는데 
		HardUser.UserInfo.Name, //UserInfo User 구조체 안에있는 네임
		HardUser.UserInfo.ID,
		HardUser.UserInfo.Password,
		HardUser.Level,
		HardUser.Age,
	)

	/////////////////////////////////////////////
	//구조체 필드명 생략한 방법

	var HardUser1 = HardUser1{User{"하하하하", "바쁘다", 200},10,20}	
	fmt.Printf("과금 유저 이름 : %s 아이디 : %s 비밀번호 : %d 레벨 : %d 나이 : %d", 
	//필드명을 생략하면 점을 한번만 찍어도 가능하다.
		HardUser1.Name, //User 구조체 안에있는 네임
		HardUser1.ID,
		HardUser1.Password,
		HardUser1.Level,
		HardUser1.Age,
	)
}
	// var Student_name string
	// var Student_age int
	// var Student_phonNumber int

	// 	Name string
	// 	Age int
	// 	score float32
	// 	Number int

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////


/*

객체지향언어다 

클래스라는 문법자체는 없지만 흉내는 낼수 있다 c++ 도 절차지향적으로 짜면은 절차지향의 언어다
확장자만 C++한다해서 객체지향언어가 아니다.

*/
