package main

import (
	"container/list"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////
type Element struct {
	Value interface{} //데이터를 저장
	Next  *Element    //다음 요소의 주소를 저장
	Prev  *Element    //이전 요소의 주소를 저장
}
/////////////////////////////////////////////////////////////////////////////
type Queue struct{
	v *list.List
}

func (q* Queue)push(val interface{}){
	q.v.PushBack(val)
}
//대기열 은행은행
func (q* Queue) Pop()interface{}{
	front := q.v.Front()
	if front != nil{
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue()*Queue{
	return &Queue{list.New()}
}

func main() {
	//리스트(List) : 각 데이터를 담고 있는 요소들을 포인터로 연결한 자료구조

	// v := list.New() //리스트 생성
	// e4:=v.PushBack(4) //뒤에 추가
	// e1:=v.PushFront(1) //앞에 추가

	// v.InsertBefore(3,e4) //
	// v.InsertAfter(2,e1) //

	// for e := v.Front(); e != nil; e = e.Next() { //정방향
	// 	fmt.Print(e.Value,"->")
	// }
	// fmt.Println()
	// for e := v.Back(); e != nil; e = e.Prev() { //역방향
	// 	fmt.Print(e.Value,"->")
	// }
	
	/*
	배열이랑 List는 여러 Data를 저장할 수 있는 자료구조 
	스텍 큐 트리구조들을 만들수가 있다. 가장 스텐다드하다고 할수 있다.
	배열이 어떻게 추가 되는지는 얘기를 해줬다.

	List에서는 각 요소가 포인터로 연결이 되어있다.
	*/

	//////////////////////////////////////////////////////////////////////////////////

		queue:=NewQueue()//큐생성
		for i := 0; i < 5; i++ {
			queue.push(i)
		}
		//때면서 출력을해준다? //pop메서드는 멘앞에 녀석을 빼주고 
		v := queue.Pop()
		for v != nil {
			fmt.Printf("%v->", v)
			v=queue.Pop()
		}
}


// /*
// 큐 Javascript로 했는데
// 큐 : 피포구조 먼저 입력한녀섴이 먼저나온다. 응가응강


// */