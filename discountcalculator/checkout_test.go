package discountcalculator

import (
	"fmt"
	"testing"
)

func TestCheckoutString(t *testing.T) {
	var tests = []struct {
		input    checkoutCode
		expected string
	}{
		{input: STANDARD_CHECKOUT, expected: "STANDARD_CHECKOUT"},
		{input: EXPRESS_CHECKOUT, expected: "EXPRESS_CHECKOUT"},
	}

	for _, test := range tests {
		if actual := fmt.Sprintf("%s", test.input); actual != test.expected {
			t.Errorf("Expected checkout code to be %s, but get %s\n", test.expected, actual)
		}

	}
}
