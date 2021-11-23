package main

import "fmt"

func main() {
	fmt.Println("13 == 13 : ", 13 == 13)
	// 13 == 13 :  true
	fmt.Println("13 == 23 : ", 13 == 23)
	// 13 == 23 :  false
	fmt.Println("13 != 13 : ", 13 != 13)
	// 13 != 13 :  false
	fmt.Println("3 != 5 : ", 3 != 5)
	// 3 != 5 :  true
	fmt.Println("0 < 1 : ", 0 < 1)
	// 0 < 1 :  true
	fmt.Println("0 > 1 : ", 0 > 1)
	// 0 > 1 :  false
	fmt.Println("0 >= 1 : ", 0 >= 1)
	// 0 >= 1 :  false
	fmt.Println("0 <= 1 : ", 0 <= 1)
	// 0 <= 1 :  true
}

