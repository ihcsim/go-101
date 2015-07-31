package stack

import "testing"

var stack Stack

func setUp() {
	stack = New()
}

func TestNew(t *testing.T) {
	setUp()

	if stack == nil {
		t.Error("Failed to create stack")
	}
}

func TestWhenPushOneItem_ThenPopThatOneItem(t *testing.T) {
	setUp()

	var pushPopTests = []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 3, expected: 3},
	}
	for _, test := range pushPopTests {
		stack.Push(test.expected)
		actual, err := stack.Pop()

		if err != nil {
			t.Errorf("Expected pop to completed without errors")
		}

		if actual != test.expected {
			t.Errorf("Expected pop to return %d, but get %d\n", test.expected, actual)
		}
	}
}

func TestWhenPushMultipleItems_ThenPopInReverseOrder(t *testing.T) {
	setUp()

	var pushPopTests = []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2}, expected: []int{2, 1}},
		{input: []int{1, 2, 3}, expected: []int{3, 2, 1}},
		{input: []int{1, 2, 3, 4}, expected: []int{4, 3, 2, 1}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1}},
	}

	for _, test := range pushPopTests {
		stack.Push(test.input...)
		for _, ex := range test.expected {
			actual, err := stack.Pop()

			if err != nil {
				t.Errorf("Expected pop to completed without errors")
			}

			if actual != ex {
				t.Errorf("Expected pop to return %d, but get %d\n", ex, actual)
			}
		}
	}
}

func TestWhenPop_ElementCountsDecrements(t *testing.T) {
	setUp()

	stack.Push(1, 2, 3, 4, 5)
	expected_length := 5
	if stack.Count() != expected_length {
		t.Errorf("Expected stack length to be %d, but get %d\n", expected_length, stack.Count())
	}

	stack.Pop()
	expected_length = 4
	if stack.Count() != expected_length {
		t.Errorf("Expected stack length to be %d, but get %d\n", expected_length, stack.Count())
	}

	stack.Pop()
	expected_length = 3
	if stack.Count() != expected_length {
		t.Errorf("Expected stack length to be %d, but get %d\n", expected_length, stack.Count())
	}

	stack.Pop()
	expected_length = 2
	if stack.Count() != expected_length {
		t.Errorf("Expected stack length to be %d, but get %d\n", expected_length, stack.Count())
	}
}

func TestWhenPushMultipleElements_ThenTopReturnsTheLastElement(t *testing.T) {
	setUp()

	var tests = []struct {
		input    []int
		expected int
	}{
		{input: []int{1}, expected: 1},
		{input: []int{1, 2}, expected: 2},
		{input: []int{1, 2, 3}, expected: 3},
		{input: []int{1, 2, 3, 4}, expected: 4},
		{input: []int{1, 2, 3, 4, 5}, expected: 5},
	}

	for _, test := range tests {
		stack.Push(test.input...)
		if stack.Top() != test.expected {
			t.Errorf("Expected top to return %d, but get %d\n", stack.Top(), test.expected)
		}
	}
}

func TestWhenTop_StackSizeDoesNotChange(t *testing.T) {
	setUp()

	var tests = []struct {
		input         []int
		expected_size int
	}{
		{input: []int{1}, expected_size: 1},
		{input: []int{1, 2}, expected_size: 2},
		{input: []int{1, 2, 3}, expected_size: 3},
		{input: []int{1, 2, 3, 4}, expected_size: 4},
		{input: []int{1, 2, 3, 4, 5}, expected_size: 5},
	}

	for _, test := range tests {
		stack.Push(test.input...)
		stack.Top()
		if stack.Count() != test.expected_size {
			t.Errorf("Expected stack size to be %d, but get %d\n", stack.Count(), test.expected_size)
		}
		stack.Clear()
	}
}

func TestWhenGivenEmptyStack_ThenPopReturnsAnError(t *testing.T) {
	setUp()

	if _, err := stack.Pop(); err == nil {
		t.Errorf("Expected an empty stack to return an error when popped")
	}
}

func TestWhenGivenEmptyStack_EmptyReturnsTrue(t *testing.T) {
	setUp()

	if !stack.Empty() {
		t.Errorf("Expected stack to be empty")
	}

	stack.Push(1)
	if stack.Empty() {
		t.Errorf("Expected stack to not be empty after a push")
	}
}
