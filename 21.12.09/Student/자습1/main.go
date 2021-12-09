package main

import "fmt"

func main() {
	//따라서 다음과 같이 fmt.Print 함수로 값을 여러 번 출력해도 모든 값이 붙어서 출력됩니다.
	fmt.Print(1)
	fmt.Print(4.5)
	fmt.Print("Hello, world!")
	fmt.Print(true)
}

/*
fmt.Print 함수는 fmt.Println 함수와 동일하지만 새 줄로 넘어가지 않고 값을 그 자리에 출력합니다.
*/