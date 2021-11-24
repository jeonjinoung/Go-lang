package main

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////

func HasRichFriend() bool {
	return true
}

func GetFriendCount () int {
	return 3
}

////////////////////////////////////////////////////////////////////////////////////////////

//멀티 반환함수 복수의 리턴값을 가지고있따.
func GetMyAge()(int,bool)  {
	return 33, true
}

func main() {
	
	color:="red"


	if color=="red"{
		fmt.Println("나는 빨간색")
	}else{
		fmt.Println("아무것도아님")
	}

	var age int = 22

	if age >= 10 && age <= 15{
		fmt.Println("dddd")
	}

	price := 35000

	if price > 50000{
		if HasRichFriend(){
			fmt.Println("오잉 난 돈이 업어")
		}else{
			fmt.Println("뿐빠이")
		}
	}else if price >= 30000 && price <=50000{
		if GetFriendCount()>3{
			fmt.Println("크음")
		}else{
			fmt.Println("제발 내자")
		}
	}else{
		fmt.Println("오느ㅜㄹ 내지갑털어봐")
	}




	////////////////////////////////////////////////////////////////////////////////////////////


	// 검사에 사용할 변수를 초기화 할때 사용
	// int인티져 , bool불 => 초기값
	//초기문 (if age, ok := getMyage())
	
	//초기문은 어떤 함수를 실행하고 결과를(반환값에 따라)검사 할때 사용한다.

	//ok && age<20 => 조건문
	//if age,ok := GetMyAge(); ok && age<20{
	
	if age,ok := GetMyAge(); ok && age<20{
		fmt.Println("젊음",age)
	}else if ok && age<30{
		fmt.Println("괜찮음",age)
	}else if ok {
		fmt.Println("굳굳굳",age)
	}else {
		fmt.Println("아무것도 아님")
	}
	fmt.Println("ㅇㅇㅇ", age)

	////////////////////////////////////////////////////////////////////////////////////////////
}
/*
	if 초기문; 조건문{
	ㄴㅇ마ㅣ릉니말ㅇㄴ}
*/
