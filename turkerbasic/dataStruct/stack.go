package dataStruct

type Stack struct {
	ll *LinkedList
}

// 새로운 스택을 만들어서 메모리 주소를 반환한다.
func NewStack() *Stack {
	return &Stack{&LinkedList{}}
}

func (s *Stack) Empty() bool {
	return s.ll.Empty()
}

func (s *Stack) Push(val int) {
	s.ll.AddNode(val)
}

func (s *Stack) Pop() int {
	back := s.ll.Back()
	s.ll.PopBack()
	return back
}
