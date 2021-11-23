package main

import "fmt"

func main() {
	a := 2
	var num int
	
	num = a
	// num = a : 2
	fmt.Println("num = a :", num)
	num += 4
	// num += 4 : 6
	// a = 2 -> num +=4 = 6
	fmt.Println("num += 4 :", num)
	num -= 2
	// num -= 2 : 4
	// a = 2 / num +=4 / num -= 2 => 6 - 2 -> 4
	fmt.Println("num -= 2 :", num)
	num *= 5
	// num *= 5 : 20
	// num의 값을 먼저 빼고 2를대입하고 a=2 그리고 num*=5 곱해주면 = 20
	fmt.Println("num *= 5 :", num)
	num /= 2
	// num /= 2 : 10  
	//num 2/=2 1 이미값음 10이고 1로 나눠봐야 10임
	
	fmt.Println("num /= 2 :", num)
	num %= 3
	// num %= 3 : 1
	fmt.Println("num %= 3 :", num)
	
	num = 3  //00000011
	num &= 2 //00000010
	fmt.Printf("num &= 2 : %08b, %d\n", num, num)
	// num &= 2 : 00000010, 2
	num |= 5 //00000101
	fmt.Printf("num |= 5 : %08b, %d\n", num, num)
	num ^= 4 //00000100
	// num |= 5 : 00000111, 7
	fmt.Printf("num ^= 4 : %08b, %d\n", num, num)
	num &^= 2 //00000010
	// num ^= 4 : 00000011, 3
	fmt.Printf("num &^= 2 : %08b, %d\n", num, num)
	num <<= 9 //00001001
	// num &^= 2 : 00000001, 1
	fmt.Printf("num <<= 9 : %08b, %d\n", num, num)
	num >>= 8 //00001000
	// num <<= 9 : 1000000000, 512
	fmt.Printf("num >>= 8 : %08b, %d\n", num, num)
	// num >>= 8 : 00000010, 2
}