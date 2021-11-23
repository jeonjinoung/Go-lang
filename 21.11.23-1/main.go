package main

import "fmt"

func main() {
	num1, num2 := 17, 5
	num3, num4 := 200, 20
	str1, str2 := "Hello", "goorm!"

	fmt.Println("num1 + num2 =", num1+num2)
	//17+5=22
	fmt.Println("str1 + str2 =", str1+str2)
	//Hellogoorm
	fmt.Println("num1 - num2 =", num1-num2)
	//17-5=12
	fmt.Println("num1 * num2 =", num1*num2)
	//17*5=85
	fmt.Println("num1 / num2 =", num1/num2)
	//17/5=3.4
	fmt.Println("num1 % num2 =", num1%num2)
	//17 % 5 = 2
	fmt.Println("num3 % num4 =", num3%num4)
	//200 % 20 = 0
}