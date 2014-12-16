package Stack

import "errors"

const max int = 10

type Stack struct {
	arr []int
	top int
}

func NewStack() *Stack {
	var s Stack = Stack{}
	s.arr = make([]int, 1, max)
	s.top = -1

	return &s
}

func (s *Stack) Push(n int) error {
	s.top++

	if len(s.arr) >= cap(s.arr) {
		s.top--
		return errors.New("Stack is full.")
	} else {
		s.arr = append(s.arr, n)
	}

	return nil
}

func (s *Stack) Pop() (int, error) {
	num := -1

	if s.top >= 0 {
		num = s.arr[s.top]
		s.top--
	} else {
		return -1, errors.New("Stack is empty")
	}

	return num, nil
}
