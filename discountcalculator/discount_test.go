package discountcalculator

import (
	"fmt"
	"testing"
)

func TestDiscountString(t *testing.T) {
	var tests = []struct {
		input    discountCode
		expected string
	}{
		{input: STANDARD_DISCOUNT, expected: "STANDARD_DISCOUNT"},
		{input: SILVER_DISCOUNT, expected: "SILVER_DISCOUNT"},
		{input: GOLD_DISCOUNT, expected: "GOLD_DISCOUNT"},
		{input: PREMIUM_DISCOUNT, expected: "PREMIUM_DISCOUNT"},
		{input: BIRTHDAY_DISCOUNT, expected: "BIRTHDAY_DISCOUNT"},
	}

	for _, test := range tests {
		if actual := fmt.Sprintf("%s", test.input); actual != test.expected {
			t.Errorf("Expected discount code to be %s, but get %s\n", test.expected, actual)
		}
	}
}
