package stack

import "errors"

type Stack []int

func New() Stack {
	var stack = []int{}
	return stack
}

func (s *Stack) Push(inputs ...int) {
	for _, item := range inputs {
		*s = append(*s, item)
	}
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return -1, errors.New("Unable to pop from an empty stack.")
	}

	head := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return head, nil
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Count() int {
	return len(*s)
}

func (s *Stack) Clear() {
	(*s) = New()
}
