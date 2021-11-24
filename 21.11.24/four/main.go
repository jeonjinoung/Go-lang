package main

import "fmt"

type MoveType int
const (
	None MoveType = iota
	North
	West
	East
	South
)

func GetDirection(angle MoveType) string {
	switch angle {
	case West:
		// fmt.Println("West")
		return  "West > 0 && West <= 45"
	case East:
		// fmt.Println("East")
		return "East > 45 && East <= 100:"
	case South:
		// fmt.Println("South")	
		return "South > 100 && South <= 200:"
	case North:
		// fmt.Println("North")
		return "North > 300:"
	default:
		return "undefinded"
	}
}

func Direction() MoveType{
	return 200
}

func main() {
	fmt.Println("방향",GetDirection(Direction()))
}

//함수이름은 GetDirction
//함수매개변수는 angle float64받음
//함수 결과 Direction타입 반환

//angle이 300이상 north반환
//angle이 0이상 45보다 작으면 west반환
//angle이 45이상 100보다 작으면 east반환
//angle이 100이상 200보다 작으면 south반환
//모든 조건이 만족하지않으면 none으로 반환