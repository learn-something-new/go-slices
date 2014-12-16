package Stack

import "errors"

const max int = 10

type Stack struct {
	arr [max]int
	top int
}

func NewStack() *Stack {
	var s Stack = Stack{}
	s.top = -1

	return &s
}

func (s *Stack) Push(n int) error {
	s.top++

	if (s.top) < max {
		s.arr[s.top] = n
	} else {
		s.top--
		return errors.New("Stack is full.")
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
