package discountcalculator

import "testing"

func TestRateFor_WhenGivenCustomerType_ThenReturnsTheCorrectDiscountRateAndType(t *testing.T) {
	var tests = []struct {
		customerType         int
		expectedRate         float64
		expectedDiscountType int
	}{
		{customerType: STANDARD, expectedRate: 0.1, expectedDiscountType: STANDARD_DISCOUNT},
		{customerType: SILVER, expectedRate: 0.15, expectedDiscountType: SILVER_DISCOUNT},
		{customerType: GOLD, expectedRate: 0.2, expectedDiscountType: GOLD_DISCOUNT},
		{customerType: PREMIUM, expectedRate: 0.25, expectedDiscountType: PREMIUM_DISCOUNT},
	}

	discountCalculator := New()
	for _, test := range tests {
		customer := NewCustomer(test.customerType)
		if rate := discountCalculator.RateFor(customer); rate != test.expectedRate {
			t.Errorf("Expected discount rate to be %.2f, but get %.2f\n", test.expectedRate, rate)
		}
		if discountType := discountCalculator.strategyCode; discountType != test.expectedDiscountType {
			t.Errorf("Expected calculation strategy to be %d, but get %d\n", test.expectedDiscountType, discountType)
		}
	}
}
