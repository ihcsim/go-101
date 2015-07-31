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
