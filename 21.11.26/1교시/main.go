package main

import "fmt"

/*

문자열 : 문자의 집함. Go에서는 UTF-8문자코드를 쓴다.
UTF-8은 영문자, 숫자 특수문자 1byte / 다른문자는 2~3byte를사용
유니코드 한 문자를 나타내는데 1~3byte크기를 사용
한글이나 한자 사용에 문제가 없다.
아스키코드와 1:1 매칭이 된다.

*/

func main() {
	//\t : tep 텝의 역할 / \n 특수문자 인식
	var str string = "Hello\t 'World\n" //더블쿼터로 묶으면 특수문자 발동
	var str1 string = `Hello\t 'World\n` //백팁으로 묶으면 특수문자 인식이안됨
	fmt.Println(str)
	fmt.Println(str1)

	//문자 하나를 표현할때는 rune타입을 사용한다.
	//go에서는 3byte의 정수타입을 제공하지 않는다.
	//rune은 int32타입의 별칭이다.(이름만 다를뿐 같은거다.)
	
	//문자하나를 쓸때는 싱글쿼터를써야함
	//var char rune = "한"
	var char rune = '한'

	//%T:Type을 볼때 사용하는것이다.
	fmt.Printf("%T\n",char)//변수타입(int32라고 나옴)
	fmt.Println(char)//타입 자체가 int32라서 (54620라고 나옴)
	//%c를 써주면 문자가 나옴??!!와우
	fmt.Printf("%c\n",char)

	var string1 string = "경일게임"
	var string2 string = "abcd"
	
	//len : 길이보는법 (객체)
	fmt.Println("string1의 길이", len(string1))// 길이 12 /한글은 3byte 씩 먹어서 3*4 = 12
	fmt.Println("string1의 길이", len(string2))// 길이 4 /영어는 1byte 씩 먹는다. 1*4 = 4
	
	//다다음주 고에서 지원하는 슬라이스라는 자료구조
	//var runes =                    []<=슬라이스라는녀섴      rune
	//슬라이스라는녀섴은 길이가 변할수 있다. 문법이다. 가변적이고 변할 수 있다.

	//string과 rune 슬라이스 타입인 []rune타입은 상호타입 면환이 가능하다.
	//슬라이스는 길이가 변할수 있는 배열(가변적이다.)

	var string3 string = "Hello World"
	//아스키코드를 사용한 문자열
	//H = 72 e= 101 l = 108 o = 111 sp(space) = 32 / W = 87 r = 114 d = 100
	var runes = []rune {72,101,108,108,111,32,87,111,114,108,100}
	
	fmt.Println(string3)
	fmt.Println(string(runes))//타입 변환을 할 경우 rune배열의 각요소에은 문자열의 각 글자가 대입이 된다.
	//이걸 이용해서 문자열의 글자수를 알 수 있다.

	var string4 string = "Hello 월드"
	var runes1  = []rune(string4)//위에 선언한 string4를 []rune타입으로 변환
	fmt.Println("string4의 길이",len(string4)) // 길이 12 // 문자열의 바이트 길이가 반환된다.
	fmt.Println("runes1의 길이",len(runes1)) //길이 8 // 각 글자들로 이루어진 배열로 반환 H1 e1 l1 l1 o1 sp1 월1 드1

	




}

/*

1교시 제목 문자열
len() : 문자길이보는 함수

말 그대로 문자 집합니다.

기본문자코드가 한글이나 한자쓰는데 문제없고

1~3byte의 크기를 먹는다. 한글자당??

Go에서는 incoding 방식을 utf-8방식을 쓴다 유니코드의 일종
(유니버셜 ...)

오늘은 문자열을 좀볼게

*/