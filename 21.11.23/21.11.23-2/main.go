package main

import "fmt"

func printScore(name string, math int, eng int, science int) {
	total := math + eng + science
	avg := total / 3
	fmt.Println(name, "의 평균 점수", avg, "이다")
}

// func result(name string, a int, b int, c int){
// 	total := a + b + c
// 	avg := total / 3
// 	if avg < 60 {
// 		fmt.Println(name, "의 평균 점수", avg, "낙제")
// 	}else if avg <= 70 {
// 		fmt.Println(name, "의 평균 점수", avg, "D")
// 	}else if avg <= 80 {
// 		fmt.Println(name, "의 평균 점수", avg, "C")
// 	}else if avg <= 90 {
// 		fmt.Println(name, "의 평균 점수", avg, "B")
// 	}else if avg <= 100{
// 		fmt.Println(name, "의 평균점수 ", avg, "A")
// 	}

// }


func main() {

//////////////////////////////////////////////////////////////////////////////////

	printScore("홍길동", 80, 80, 100)
	printScore("배추도사", 100, 80, 100)
	printScore("은비까비", 80, 70, 100)
	
//////////////////////////////////////////////////////////////////////////////////////

	// result("전진영", 60,50,55)
	// result("전주영", 65,75,60)
	// result("김복인", 80,70,85)
	// result("전용식", 85,80,85)
	// result("전용식", 90,95,90)

}

// func grade() int {
// 	printScore(avg) <=100, =>90 = string "A"	
// 	printScore(avg) <=80, =>89 = string "B"
// 	printScore(avg) <=79, =>79 = string "C"
// 	printScore(avg) <=69, =>60 = string "D"
// 	printScore(avg) <=59 = string "낙제"
// }



//90~100점이면 A등급
//80~89이면 B등급
//70~79아면 C등급
//60~69이면 D등급
//59이하면 낙제
