package main

import "fmt"

/*
	 1 func 2 Add ( 3 a int, b int) 4 int {

		5 return a + b

	}
	1 func 함수키워드
	2 Add 함수이름
	3 a int, b int 매개변수 객체 파라미터
	4 반환타입 -> int
	5 함수의코드블럭

*/  
func Add(a int, b int) int {
	return a + b
}

func printScore(name string, math int, eng int, science int){
	total:=math+eng+science
	avg:=total/3
	fmt.Println(name,"의 평균 점수", avg,"이다")
}

//멀티반환함수 : 여러개를 리턴할수 있다. 여러개 리턴값을 가지고있따.
/*
func Divide(a int, b int) =  (a,b int){
매개변수 타입이 같으면 아래와 같이 생략이 가능하다	
}
타입이 여러개면 소괄호로 묶어준다. (int,bool)
*/

// func Divide(a,b int)(int,bool)  {
// 	if b==0{
// 		//(int = 0,bool = false)

// 		return 0, false
// 	}
// 	//int = a / b bool = true
// 	return a / b, true
	
// }

//변수명을 이용한 반홤
//해당값을 명시적으로 반환하지않아도 값을 리턴할 수 있다.
//그냥 리턴만 하는거니까
func Divide(a,b int)(result int, success bool)  {
	if b==0{
		//result int  = result = 0
		result = 0
		//success bool sucess = false
		success = false
		return // 명시적으로 반환할 값을 지정하지 않음
	}
	result = a / b
	success = true
	return 
}



func main() {
	c := Add(3, 6)//첫번째 함수호출
	fmt.Println(c)

	printScore("홍길동", 80, 80, 100)
	printScore("배추도사", 100, 80, 100)
	printScore("은비까비", 80, 70, 100)


	a, success :=Divide(9,3)
	fmt.Println(a,success)//a = 3  success = true
	//fmt.Println(a,success)// a/b 리턴값이 a에 들어오고 true의 값이 sucess로 들어와서 true로 나오게된다.
	
	/*

	3하고 true가 나왔는데 여기에 지금 
	
	*/
	



}
/*




*/


/*

//함수를 쓰는 이유
//가속성좋아지고 편하고 등등 개좋다

함수로 호출하면 이값 넣으면 이값나오네!!

보낸값을 그대로 쓰는게 아니다 3하고 6을 넣었다고해서 그값을 그대로 쓰는게 아니다.

함수를 호출하면

인수는  3,6 아규먼트 이 인수들은 매개변수의 매개변수에 복사가 된다.
3과 6일 3이 a에 복사 6이 b에 복사가 된다 = return a + b

매개변수랑 함수내에서 선언된 변수들은 함수가 종료가되는순간 변수의 범위를 벗어난다.
호출함 함수가 종료가 되면 함수에서 사용한 지역변수에 접근이 안된다. 
로컬변수는 당연히 함수가 호출이 끝나면 모든 지역변수들은 날라가게 되있다.

그다음에 fmt.Println(c)는 9가 나온다.const


*/


//함수를 이용해서 평균값을 구해서
//학점뽑기

//90~100점이면 A등급
//80~89이면 B등급
//70~79아면 C등급
//60~69이면 D등급
//59이하면 낙제
