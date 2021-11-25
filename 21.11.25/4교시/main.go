package main

import "fmt"

//구조체를 이용해서 원하는 데이터를 출력해보세요 구조체를 이용하세요!!
//영화가 됐던 뭐가됐던 뽑아봐라!!

type Student struct {
	Name    string
	Age     int
	Number  int
	Address string
}

type HardUser struct {
	ID       string
	Password int
}

type User struct {
	Studentinfo  Student
	HardUserinfo HardUser
	Hooby        string
	Time         int
}

func main() {

	var myStory = User{Student{"전진영", 29, 75673917, "감일동"}, HardUser{"mbctg", 12051288}, "춤", 25}
	fmt.Println(myStory)
}

//나만의 장기자랑 만들기


