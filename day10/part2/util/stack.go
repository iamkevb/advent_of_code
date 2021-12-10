package util

import "fmt"

type Stack struct {
	elem []rune
}

func (s *Stack) Push(r rune) {
	s.elem = append(s.elem, r)
}

func (s *Stack) Pop() {
	if s.Empty() {
		panic("Cannot remove from empty stack")
	}
	s.elem = s.elem[:len(s.elem)-1]
}

func (s *Stack) Top() rune {
	if s.Empty() {
		panic("Cannot remove from empty stack")
	}
	return s.elem[len(s.elem)-1]
}

func (s *Stack) Empty() bool {
	return len(s.elem) == 0
}

func (s *Stack) Print() {
	fmt.Println(string(s.elem))
}

func (s *Stack) String() string {
	return string(s.elem)
}

func NewStack() *Stack {
	return &Stack{
		elem: make([]rune, 0),
	}
}
