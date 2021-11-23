package main

import "fmt"

func main() {
	var a bool = true
	b := false

	fmt.Println("0 && 0 : ", b && b)
	// 0 && 0 :  false
	fmt.Println("0 && 1 : ", b && a)
	// 0 && 1 :  false
	fmt.Println("1 && 1 : ", a && a)
	// 1 && 1 :  true
	fmt.Println("0 || 0 : ", b || b)
	// 0 || 0 :  false
	fmt.Println("0 || 1 : ", b || a)
	// 0 || 1 :  true
	fmt.Println("1 || 1 : ", a || a)
	// 1 || 1 :  true
	fmt.Println("!1 ", !true)
	// !1  false
	fmt.Println("!0 ", !false)
	// !0  true
}

