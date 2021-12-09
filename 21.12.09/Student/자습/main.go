package main

import "fmt"

type Data struct{ a, b int }

func main() {
	
	var num1 int = 10
	var num2 float32 = 3.2
	var num3 complex64 = 2.5 + 8.1i
	var s string = "Hello, world!"
	var b bool = true
	var a []int = []int{1, 2, 3}
	var m map[string]int = map[string]int{"Hello": 1}
	var p *int = new(int)
	var data Data = Data{1, 2}
	var i interface{} = 1

	fmt.Println(num1) // 10: 정수 출력
	fmt.Println(num2) // 3.2: 실수 출력
	fmt.Println(num3) // (2.5+8.1i): 복소수 출력
	fmt.Println(s)    // Hello, world!: 문자열 출력
	fmt.Println(b)    // true: 불 출력
	fmt.Println(a)    // [1 2 3]: 슬라이스 출력
	fmt.Println(m)    // map[Hello:1]: 맵 출력
	fmt.Println(p)    // 0xc0820062d0: 포인터(메모리 주소) 출력
	fmt.Println(data) // {1 2}: 구조체 출력
	fmt.Println(i)    // 1: 인터페이스 출력

}
/*
1. func Print(a ...interface{}) (n int, err error): 값을 그 자리에 출력(새 줄로 넘어가지 않음)

2. func Println(a ...interface{}) (n int, err error): 값을 출력한 뒤 새 줄로 넘어감(개행)

3. func Printf(format string, a ...interface{}) (n int, err error): 형식을 지정하여 값을 출력
*/