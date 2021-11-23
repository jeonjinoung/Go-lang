package main

import "fmt"

func main() {
	count1, count2 := 1, 10.4
	
	count1++
	count2--
	
	fmt.Println("count1++ :", count1)
	//count1++ : 2
	fmt.Println("count2-- :", count2)
	//count2-- : 9.4
}
