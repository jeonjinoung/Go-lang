package main

import "fmt"

//별명(별칭) => ColorType / int가 붙음으로인한
type ColorType int
const (
	//같은패턴을 반복하는것이니까 생략가능하다.
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"		
	default:
		return "undefinded"
	}
}

func GetColor() ColorType {
	return Blue
}


func main() {
	fmt.Println("색깔",colorToString(GetColor()))
}

//방향을 나타내는것??

//열거형과 사용하면 효력이 좋다.
//const 열거형에 따라 로직을 변경할때 쓰면 유횽하다

/*

어렵지 않아오 다들 할 수 있어용!!

*/