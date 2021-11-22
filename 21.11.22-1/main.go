package main

import "fmt"

func main()  {
// 	fmt.Println("Hellow World")

//   //var a<변수명> int<type> = <초기값>
// 	var a int
// 	var f float32 = 11.
// 	var i,j, k int = 1, 2, 3
	
// 	var d = 5

// 	//type을 생략하는 방법
// 	//선언대입문 : :=를 써서 var키워드, type을 생략
// 	//Short Assignmetn Statement (:=)를 사용할 수 도 있다. 근데 이건 지역변수 선언할 때만 가능하다. 
// 	num2 :=10
// 	fmt.Println("키워드와 타입생략 방법",num2)


// 	fmt.Println(a)
// 	fmt.Println(f)
// 	fmt.Println(i,j,k)
// 	fmt.Println(d)


//////////////////////////////////////////////////////////////////////////////////////////////////////////

	/*
	const 키워드를 사용하여 상수를 선언한다. 걍 맨 앞에  var 대신 const 붙이면 됨.
	여러개 묶어서 한번에 const 지정할 수 있음.
	*/

	const hi = "Hi"

	const (
    Sky = "Blue"
    Rose = "Red"
    Gyuri = "Awesome"
	)

	fmt.Println(hi)
	fmt.Println(Sky, Rose, Gyuri)

//////////////////////////////////////////////////////////////////////////////////////////////////////////

	// Type Conversion
	// Go에서는 암묵적인 type conversion이 이뤄지지 않는다. 항상 명시적으로 지정해 주어야 한다.
	
	// Type(var)와 같이 표현한다.
	

    var i int = 100
    var u uint = uint(i)
    var f float32 = float32(i)  
    fmt.Println(i,f, u)
 
    str := "ABC"
    bytes := []byte(str)
    str2 := string(bytes)
    println(bytes, str2)


}




