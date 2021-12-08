package main

import (
	"fmt"
	"os"
)

//인수타입앞에 (...) 해당 타입 인수를 여거개 받는 가변인수
// 가변인수는 함수내부에서 해당 type의 slice처리가 된다.
func sum(nums...int)int{
	sum:=0

	fmt.Printf("type : %T\n", nums)

	for _, v := range nums {
		sum+= v
	}
	return sum
}

//사용안함
// func Print(args ...interface{})string{//모든 타입을 받는 가변인수
// 	for _, arg := range args { //모든 인수를 돌면서
// 		switch b:=arg.(type){ //인수타입에 따라서
// 		case bool:
// 			val := arg.(bool) //변환
// 		case int:
// 			val := arg.(int) //변환
// 		}
// 	}
// }

func main() {

	/*실행함
	fmt. Println()
	fmt. Println(1)
	fmt. Println(1,2,3,4,5,6)
	*/

	// 실행함 위의 sum slice타입
	fmt.Println(sum(1,2,3,4,5,6,7))
	fmt.Println(sum(1,2))
	fmt.Println(sum())
	
	///////////////////////////////////////////////

	//가변인수를 빈 인터페이스로 받아주면된다.
	fmt.Println(1, "ㅇㅇㅇ", 1.21321)

	//해당함수가 실행하기전에 
	//defer 명령
	//defer라는 녀석은 지연을 시켜주는 녀섴
	///////////////////////////////////////////////////
	//실행함
	//os페케지에서 사용하는 함수를 썼다.
	//파일이 생성되고 

	f, err := os.Create("text.txt")
	if err != nil {

		fmt.Println("실패")

		return
	}

	//역순으로 진행된다.
	defer fmt.Println("호출")//지연
	defer f.Close()//지연
	defer fmt.Println("파일 닫음")//지연
	fmt.Println("작성")
	fmt.Fprintln(f,"Hello World")
}
/*

없을수도 있고 한게일수도있고 여러개일수도 있다.
가변인수함수라고한다. 가변적이다 바뀔수있다. -> C++에서는 이거를  함수 오버로딩이라고 한다.

가변인수는 함수내부에서 해당 type의 slice처리가 된다.

*/