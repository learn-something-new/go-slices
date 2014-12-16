package Stack

import "errors"

const max int = 10

type Stack struct {
	arr []int
	top int
}

func NewStack() *Stack {
	var s Stack = Stack{}
	s.arr = make([]int, max)
	s.top = -1

	return &s
}

func (s *Stack) Top() int {
	return s.top
}

func (s *Stack) Len() int {
	return len(s.arr)
}

func (s *Stack) Cap() int {
	return cap(s.arr)
}

func (s *Stack) Push(n int) error {
	s.top++

	if (s.top) >= cap(s.arr) {
		newArr := make([]int, ((len(s.arr) + 1) * 2), ((len(s.arr) + 1) * 2))
		copy(newArr, s.arr)
		s.arr = newArr
	}

	if (s.top) >= len(s.arr) {
		s.top--
		return errors.New("Stack is full")
	}

	s.arr[s.top] = n

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
