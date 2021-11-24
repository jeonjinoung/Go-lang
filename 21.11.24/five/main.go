package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i==3 {
	// 		countinue
	// 	}
	// 	if i==6 {
	// 		break
	// 	}
	// }

	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	dan :=2
	b:= 1
	for {
		for {
		fmt.Printf("%d*%d = %d\n", dan, b, dan*b)
		b++
		if b ==10{
			break
			}
		}
		b=1
		dan++
		if dan==10 {
			break
		}
	}
}

//1.초기문 생략

// for ; 조건 ; 증감{
// }

//2. 증감생략
// for i := 0; i < count;{

// }

//3. 조건문만 있는경우
// for ; 조건; {
// }

//4.무한루프( while이랑 같음)
// for true {
// }

//위아 아래와 동일한 무한루프이다.
//4.1 무한루프 트루가없어도됨
// for{
// }
