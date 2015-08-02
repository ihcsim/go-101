package stack

import "testing"

var stack Stack

func setUp() {
	stack.Clear()
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
		input    interface{}
		expected interface{}
	}{
		{input: 1, expected: 1},
		{input: 2.0, expected: 2.0},
		{input: "hello", expected: "hello"},
	}
	for _, test := range pushPopTests {
		stack.Push(test.expected)
		actual, err := stack.Pop()

		if err != nil {
			t.Errorf("Expected pop to completed without errors")
		}

		if actual != test.expected {
			t.Errorf("Expected pop to return %v, but get %v\n", test.expected, actual)
		}
	}
}

func TestWhenPushMultipleItems_ThenPopInReverseOrder(t *testing.T) {
	setUp()

	var pushPopTests = []struct {
		input    []interface{}
		expected []interface{}
	}{
		{input: []interface{}{1, 2}, expected: []interface{}{2, 1}},
		{input: []interface{}{true, true, false}, expected: []interface{}{false, true, true}},
		{input: []interface{}{'A', 'B', 'C', 'D'}, expected: []interface{}{'D', 'C', 'B', 'A'}},
		{input: []interface{}{"AB", "CD", "EF", "GH", "IJ"}, expected: []interface{}{"IJ", "GH", "EF", "CD", "AB"}},
	}

	for _, test := range pushPopTests {
		stack.Push(test.input...)
		for _, ex := range test.expected {
			actual, err := stack.Pop()

			if err != nil {
				t.Errorf("Expected pop to completed without errors")
			}

			if actual != ex {
				t.Errorf("Expected pop to return %v, but get %v\n", ex, actual)
			}
		}
	}
}

func TestWhenPop_ElementCountsDecrements(t *testing.T) {
	setUp()

	stack.Push(1, 2, 3, 4, 5)
	expected_length := 5
	if stack.Count() != expected_length {
		t.Errorf("Expected stack length to be %v, but get %v\n", expected_length, stack.Count())
	}

	stack.Pop()
	expected_length = 4
	if stack.Count() != expected_length {
		t.Errorf("Expected stack length to be %v, but get %v\n", expected_length, stack.Count())
	}

	stack.Pop()
	expected_length = 3
	if stack.Count() != expected_length {
		t.Errorf("Expected stack length to be %v, but get %v\n", expected_length, stack.Count())
	}

	stack.Pop()
	expected_length = 2
	if stack.Count() != expected_length {
		t.Errorf("Expected stack length to be %v, but get %v\n", expected_length, stack.Count())
	}
}

func TestWhenPushMultipleElements_ThenTopReturnsTheLastElement(t *testing.T) {
	setUp()

	var tests = []struct {
		input    []interface{}
		expected interface{}
	}{
		{input: []interface{}{1}, expected: 1},
		{input: []interface{}{true, false}, expected: false},
		{input: []interface{}{'A', 'B', 'C'}, expected: 'C'},
		{input: []interface{}{"AB", "CD", "EF", "GH"}, expected: "GH"},
		{input: []interface{}{1.00, 2.00, 3.00, 4.00, 5.00}, expected: 5.00},
	}

	for _, test := range tests {
		stack.Push(test.input...)
		top, _ := stack.Top()
		if top != test.expected {
			t.Errorf("Expected top to return %v, but get %v\n", top, test.expected)
		}
	}
}

func TestWhenTop_StackSizeDoesNotChange(t *testing.T) {
	setUp()

	var tests = []struct {
		input         []interface{}
		expected_size int
	}{
		{input: []interface{}{1}, expected_size: 1},
		{input: []interface{}{1.0, 2.0}, expected_size: 2},
		{input: []interface{}{"1", "2", "3"}, expected_size: 3},
		{input: []interface{}{'A', 'B', 'C', 'D'}, expected_size: 4},
		{input: []interface{}{true, false, true, false, true}, expected_size: 5},
	}

	for _, test := range tests {
		stack.Push(test.input...)
		stack.Top()
		if stack.Count() != test.expected_size {
			t.Errorf("Expected stack size to be %v, but get %v\n", stack.Count(), test.expected_size)
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

func TestWhenGivenEmptyStack_ThenTopReturnsAnError(t *testing.T) {
	setUp()

	if _, err := stack.Top(); err == nil {
		t.Errorf("Expected an empty stack to return an error when topped")
	}
}

func TestWhenNoElementsInTheStack_EmptyReturnsTrue(t *testing.T) {
	setUp()

	if !stack.Empty() {
		t.Errorf("Expected stack to be empty")
	}

	stack.Push(1)
	if stack.Empty() {
		t.Errorf("Expected stack to not be empty after a push")
	}
}
