package stack

import "errors"

type Stack []interface{}

func (s *Stack) Push(inputs ...interface{}) {
	for _, item := range inputs {
		*s = append(*s, item)
	}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return -1, errors.New("Unable to pop from an empty stack.")
	}

	head := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return head, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("Unable to top from an empty stack")
	}
	return (*s)[len(*s)-1], nil
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Count() int {
	return len(*s)
}

func (s *Stack) Clear() {
	(*s) = []interface{}{}
}
